package deviceinteractlogic

import (
	"context"
	"encoding/json"
	"gitee.com/i-Things/share/devices"
	"gitee.com/i-Things/share/domain/schema"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/service/dmsvr/internal/domain/deviceMsg/msgThing"
	"github.com/i-Things/things/service/dmsvr/internal/repo/cache"
	"time"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyLatestReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	schema *schema.Model
}

func NewGetPropertyLatestReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyLatestReplyLogic {
	return &GetPropertyLatestReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *GetPropertyLatestReplyLogic) initMsg(productID string) error {
	var err error
	l.schema, err = l.svcCtx.SchemaRepo.GetSchemaModel(l.ctx, productID)
	if err != nil {
		return errors.System.AddDetail(err)
	}
	return nil
}

// 请求设备获取设备最新属性
func (l *GetPropertyLatestReplyLogic) GetPropertyLatestReply(in *dm.GetPropertyLatestReplyReq) (*dm.GetPropertyLatestReplyResp, error) {
	l.Infof("%s req=%+v", utils.FuncName(), in)
	if err := CheckIsOnline(l.ctx, l.svcCtx, devices.Core{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}); err != nil {
		return nil, err
	}

	err := l.initMsg(in.ProductID)
	if err != nil {
		return nil, err
	}

	MsgToken := devices.GenMsgToken(l.ctx)

	req := msgThing.Req{
		CommonMsg: deviceMsg.CommonMsg{
			Method:    deviceMsg.GetReport,
			MsgToken:  MsgToken,
			Timestamp: time.Now().UnixMilli(),
		},
		Identifiers: in.DataIDs,
	}

	payload, _ := json.Marshal(req)
	reqMsg := deviceMsg.PublishMsg{
		Handle:     devices.Thing,
		Type:       msgThing.TypeProperty,
		Payload:    payload,
		Timestamp:  time.Now().UnixMilli(),
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}
	err = cache.SetDeviceMsg(l.ctx, l.svcCtx.Cache, deviceMsg.ReqMsg, &reqMsg, req.MsgToken)
	if err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.PubDev.ReqToDeviceSync(l.ctx, &reqMsg, func(payload []byte) bool {
		var dresp msgThing.Resp
		err = utils.Unmarshal(payload, &dresp)
		if err != nil { //如果是没法解析的说明不是需要的包,直接跳过即可
			return false
		}
		if dresp.MsgToken != req.MsgToken { //不是该请求的回复.跳过
			return false
		}
		return true
	})
	if err != nil {
		return nil, err
	}

	var dresp msgThing.Req
	err = utils.Unmarshal(resp, &dresp)
	if err != nil {
		return nil, err
	}
	var params []byte
	if len(dresp.Params) > 0 {
		params, _ = json.Marshal(dresp.Params)
	}
	return &dm.GetPropertyLatestReplyResp{
		MsgToken:  dresp.MsgToken,
		Msg:       dresp.Msg,
		Code:      dresp.Code,
		Timestamp: dresp.GetTimeStamp(time.Now().UnixMilli()).UnixMilli(),
		Params:    string(params),
	}, nil
}
