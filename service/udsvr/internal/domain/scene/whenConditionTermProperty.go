package scene

import (
	"context"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/domain/schema"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	devicemsg "gitee.com/unitedrhino/things/service/dmsvr/client/devicemsg"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

// TermProperty 物模型类型 属性
type TermProperty struct {
	AreaID           int64    `json:"areaID,string,omitempty"` //仅做记录
	ProductName      string   `json:"productName,omitempty"`   //产品名称,只读
	ProductID        string   `json:"productID,omitempty"`     //产品id
	DeviceName       string   `json:"deviceName,omitempty"`
	DeviceAlias      string   `json:"deviceAlias,omitempty"`
	DataID           string   `json:"dataID,omitempty"` //属性的id   aa.bb.cc
	DataName         string   `json:"dataName,omitempty"`
	SchemaAffordance string   `json:"schemaAffordance,omitempty"` //只读,返回物模型定义
	TermType         CmpType  `json:"termType,omitempty"`         //动态条件类型  eq: 相等  not:不相等  btw:在xx之间  gt: 大于  gte:大于等于 lt:小于  lte:小于等于   in:在xx值之间
	Values           []string `json:"values,omitempty"`           //条件值 参数根据动态条件类型会有多个参数
}

func (c *TermProperty) Validate(repo CheckRepo) error {
	if c == nil {
		return nil
	}
	if err := c.TermType.Validate(c.Values); err != nil {
		return err
	}
	if repo.Info.DeviceMode == DeviceModeSingle {
		c.ProductID = repo.Info.ProductID
		c.DeviceName = repo.Info.DeviceName
	}
	if c.ProductID == "" {
		return errors.Parameter.AddMsg("执行条件设备类型中的产品id需要填写")
	}
	if c.DeviceName == "" {
		return errors.Parameter.AddMsg("执行条件设备类型中的设备名需要填写")
	}
	if len(c.DataID) == 0 {
		return errors.Parameter.AddMsg("执行条件设备类型中的标识符需要填写")
	}
	if repo.Info.DeviceMode != DeviceModeSingle {
		c.DeviceAlias = GetDeviceAlias(repo.Ctx, repo.DeviceCache, c.ProductID, c.DeviceName)
	}
	v, err := repo.SchemaCache.GetData(repo.Ctx, devices.Core{ProductID: c.ProductID, DeviceName: c.DeviceName})
	if err != nil {
		return err
	}
	p := v.Property[c.DataID]
	if p == nil {
		return errors.Parameter.AddMsg("dataID不存在")
	}
	c.SchemaAffordance = schema.DoToAffordanceStr(p)
	if c.DataName == "" {
		c.DataName = p.Name
	}
	pi, err := repo.ProductCache.GetData(repo.Ctx, c.ProductID)
	if err != nil {
		return err
	}
	c.ProductName = pi.ProductName
	return nil
}
func (c *TermProperty) IsHit(ctx context.Context, columnType TermColumnType, repo CheckRepo) bool {
	sm, err := repo.SchemaCache.GetData(ctx, devices.Core{ProductID: c.ProductID, DeviceName: c.DeviceName})
	if err != nil {
		logx.WithContext(ctx).Errorf("%s.GetSchemaModel err:%v", utils.FuncName(), err)
		return false
	}
	var val string
	var dataType schema.DataType
	switch columnType {
	case TermColumnTypeProperty:
		dataID := strings.Split(c.DataID, ".")
		info, err := repo.DeviceMsg.PropertyLogLatestIndex(ctx, &devicemsg.PropertyLogLatestIndexReq{ProductID: c.ProductID, DeviceName: c.DeviceName, DataIDs: dataID[:1]})
		if err != nil {
			logx.WithContext(ctx).Errorf("%s.PropertyLatestIndex err:%v", utils.FuncName(), err)
			return false
		}
		if len(info.List) == 0 {
			logx.WithContext(ctx).Errorf("%s.PropertyLatestIndex err:dataID is not right:%v", utils.FuncName(), c.DataID[0])
			return false
		}
		if info.List[0].Timestamp != 0 { //如果有值
			dataType = sm.Property[dataID[0]].Define.Type
			def := sm.Property[dataID[0]].Define
			switch def.Type {
			case schema.DataTypeStruct:
				if len(dataID) < 2 { //必须指定到结构体的成员
					return false
				}
				var dataMap = map[string]any{}
				utils.Unmarshal([]byte(info.List[0].Value), &dataMap)
				v, ok := dataMap[dataID[1]]
				if ok {
					val = cast.ToString(v)
					dataType = def.Spec[dataID[1]].DataType.Type
				}
			case schema.DataTypeArray:
				logx.WithContext(ctx).Errorf("%s scene not support array yet", utils.FuncName())
				return false
			default:
				val = info.List[0].Value
			}
		}
		return c.TermType.IsHit(dataType, val, c.Values)
	case TermColumnTypeEvent:
		logx.WithContext(ctx).Errorf("scene not support event yet")
		return false
	}
	return true
}
