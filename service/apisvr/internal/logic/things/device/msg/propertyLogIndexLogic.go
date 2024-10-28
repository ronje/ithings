package msg

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/apisvr/internal/logic"
	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type PropertyLogIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPropertyLogIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PropertyLogIndexLogic {
	return &PropertyLogIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctxs.WithDefaultRoot(ctx),
		svcCtx: svcCtx,
	}
}

func (l *PropertyLogIndexLogic) PropertyLogIndex(req *types.DeviceMsgPropertyLogIndexReq) (resp *types.DeviceMsgPropertyIndexResp, err error) {
	dmResp, err := l.svcCtx.DeviceMsg.PropertyLogIndex(l.ctx, &dm.PropertyLogIndexReq{
		DeviceNames: req.DeviceNames,
		ProductID:   req.ProductID,
		DataID:      req.DataID,
		TimeStart:   req.TimeStart,
		TimeEnd:     req.TimeEnd,
		Interval:    req.Interval,
		ArgFunc:     req.ArgFunc,
		Fill:        req.Fill,
		Order:       req.Order,
		Page:        logic.ToDmPageRpc(req.Page),
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.GetDeviceData req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	info := make([]*types.DeviceMsgPropertyLogInfo, 0, len(dmResp.List))
	for _, v := range dmResp.List {
		info = append(info, &types.DeviceMsgPropertyLogInfo{
			Timestamp: v.Timestamp,
			DataID:    v.DataID,
			Value:     v.Value,
		})
	}
	return &types.DeviceMsgPropertyIndexResp{
		Total: dmResp.Total,
		List:  info,
	}, nil
}
