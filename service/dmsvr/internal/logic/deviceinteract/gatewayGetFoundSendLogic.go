package deviceinteractlogic

import (
	"context"
	"encoding/json"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/domain/deviceMsg"
	"gitee.com/unitedrhino/share/domain/deviceMsg/msgGateway"
	"gitee.com/unitedrhino/share/domain/deviceMsg/msgThing"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	"time"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GatewayGetFoundSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGatewayGetFoundSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GatewayGetFoundSendLogic {
	return &GatewayGetFoundSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 实时获取网关拓扑关系
func (l *GatewayGetFoundSendLogic) GatewayGetFoundSend(in *dm.GatewayGetFoundReq) (*dm.Empty, error) {
	_, err := logic.SchemaAccess(l.ctx, l.svcCtx, def.AuthRead, devices.Core{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}, nil)
	if err != nil {
		return nil, err
	}
	//return &dm.Empty{}, nil
	var protocolCode string
	if protocolCode, err = CheckIsOnline(l.ctx, l.svcCtx, devices.Core{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}); err != nil {
		return nil, err
	}
	MsgToken := devices.GenMsgToken(l.ctx, l.svcCtx.NodeID)
	req := msgGateway.Msg{
		CommonMsg: deviceMsg.CommonMsg{
			Method:   deviceMsg.GetFound,
			MsgToken: MsgToken,
			//Timestamp: time.Now().UnixMilli(),
		},
	}
	payload, _ := json.Marshal(req)
	reqMsg := deviceMsg.PublishMsg{
		Handle:       devices.Gateway,
		Type:         msgGateway.TypeTopo,
		Payload:      payload,
		Timestamp:    time.Now().UnixMilli(),
		ProductID:    in.ProductID,
		DeviceName:   in.DeviceName,
		ProtocolCode: protocolCode,
	}
	var resp []byte
	resp, err = l.svcCtx.PubDev.ReqToDeviceSync(l.ctx, &reqMsg, time.Second*10, func(payload []byte) bool {
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
	var dresp msgGateway.Msg
	err = utils.Unmarshal(resp, &dresp)
	if err != nil {
		return nil, err
	}
	if dresp.Code != errors.OK.GetCode() {
		return nil, errors.DeviceResp.AddMsgf("设备返回错误: msg:%v code:%v", dresp.Msg, dresp.Code)
	}
	return &dm.Empty{}, nil
	//if dresp.Payload != nil || len(dresp.Payload.Devices) == 0 {
	//	return &dm.GatewayGetFoundResp{}, nil
	//}
	//var devs []*devices.Core
	//for _, v := range dresp.Payload.Devices {
	//	devs = append(devs, &devices.Core{
	//		ProductID:  v.ProductID,
	//		DeviceName: v.DeviceName,
	//	})
	//}
	//diDB := relationDB.NewDeviceInfoRepo(l.ctx)
	////只获取已经入网的设备,未入网的设备需要网关自己注册或提前入网
	//dis, err := diDB.FindByFilter(l.ctx, relationDB.DeviceFilter{Cores: devs}, nil)
	//if err != nil {
	//	return nil, err
	//}
	////gdDB := relationDB.NewGatewayDeviceRepo(l.ctx)
	////gdDB.FindByFilter()
	//fmt.Println(dis)
	//return &dm.GatewayGetFoundResp{}, nil
}
