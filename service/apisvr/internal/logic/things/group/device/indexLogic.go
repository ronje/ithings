package device

import (
	"context"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/apisvr/internal/logic"
	"gitee.com/unitedrhino/things/service/apisvr/internal/logic/things"
	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.GroupDeviceIndexReq) (resp *types.GroupDeviceIndexResp, err error) {
	var list []*types.DeviceInfo
	gd, err := l.svcCtx.DeviceG.GroupDeviceIndex(l.ctx, &dm.GroupDeviceIndexReq{
		Page:       logic.ToDmPageRpc(req.Page),
		GroupID:    req.GroupID,
		ProductID:  req.ProductID,
		DeviceName: req.DeviceName,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.DeviceGroup GroupDeviceIndex req=%v err=%+v", utils.FuncName(), req, er)
		return nil, err
	}
	for _, v := range gd.List {
		pi := things.InfoToApi(l.ctx, l.svcCtx, v, things.DeviceInfoWith{Properties: req.WithProperties, Profiles: req.WithProfiles})
		list = append(list, pi)
	}
	return &types.GroupDeviceIndexResp{
		List:  list,
		Total: gd.Total,
	}, nil
}
