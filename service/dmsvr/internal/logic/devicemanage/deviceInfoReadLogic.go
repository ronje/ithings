package devicemanagelogic

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceInfoReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DiDB *relationDB.DeviceInfoRepo
}

func NewDeviceInfoReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceInfoReadLogic {
	return &DeviceInfoReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		DiDB:   relationDB.NewDeviceInfoRepo(ctx),
	}
}

// 获取设备信息详情
func (l *DeviceInfoReadLogic) DeviceInfoRead(in *dm.DeviceInfoReadReq) (*dm.DeviceInfo, error) {

	if ctxs.GetUserCtx(l.ctx).IsAdmin {
		l.ctx = ctxs.WithAllProject(l.ctx)
	}
	di, err := l.DiDB.FindOneByFilter(l.ctx,
		relationDB.DeviceFilter{ProductID: in.ProductID, DeviceNames: []string{in.DeviceName},
			SharedType: def.SelectTypeAll})
	if err != nil && !errors.Cmp(err, errors.NotFind) {
		l.Error(err)
		return nil, err
	}
	if di == nil {
		di, err = l.DiDB.FindOneByFilter(l.ctx,
			relationDB.DeviceFilter{DeviceNames: []string{in.DeviceName},
				SharedType: def.SelectTypeAll})
		if err != nil {
			l.Error(err)
			return nil, err
		}
	}
	pb := logic.ToDeviceInfo(l.ctx, l.svcCtx, di)
	if in.WithGateway && pb.DeviceType == def.DeviceTypeSubset {
		gd, err := relationDB.NewGatewayDeviceRepo(l.ctx).FindOneByFilter(l.ctx, relationDB.GatewayDeviceFilter{SubDevice: &devices.Core{
			ProductID:  di.ProductID,
			DeviceName: di.DeviceName,
		}})
		if err != nil && !errors.Cmp(err, errors.NotFind) {
			return nil, err
		}
		if gd != nil {
			ddi, err := l.svcCtx.DeviceCache.GetData(l.ctx, devices.Core{
				ProductID:  gd.GatewayProductID,
				DeviceName: gd.GatewayDeviceName,
			})
			if err != nil && !errors.Cmp(err, errors.NotFind) {
				return nil, err
			}
			if ddi != nil {
				pb.Gateway = ddi
			}
		}
	}
	return pb, nil
}
