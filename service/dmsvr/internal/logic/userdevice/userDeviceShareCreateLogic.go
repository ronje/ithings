package userdevicelogic

import (
	"context"
	"gitee.com/i-Things/core/service/syssvr/pb/sys"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/errors"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeviceShareCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeviceShareCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeviceShareCreateLogic {
	return &UserDeviceShareCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分享设备
func (l *UserDeviceShareCreateLogic) UserDeviceShareCreate(in *dm.UserDeviceShareInfo) (*dm.WithID, error) {
	di, err := relationDB.NewDeviceInfoRepo(l.ctx).FindOneByFilter(l.ctx, relationDB.DeviceFilter{ProductID: in.Device.ProductID, DeviceName: in.Device.DeviceName})
	if err != nil {
		return nil, err
	}
	pi, err := l.svcCtx.ProjectM.ProjectInfoRead(l.ctx, &sys.ProjectWithID{ProjectID: int64(di.ProjectID)})
	if err != nil {
		return nil, err
	}
	uc := ctxs.GetUserCtx(l.ctx)
	if pi.AdminUserID != uc.UserID {
		return nil, errors.Permissions.AddMsg("只有所有者才能分享设备")
	}
	po := relationDB.DmUserDeviceShare{
		UserID:     in.UserID,
		ProductID:  in.Device.ProductID,
		DeviceName: in.Device.DeviceName,
		AccessPerm: in.AccessPerm,
		SchemaPerm: in.SchemaPerm,
	}
	err = relationDB.NewUserDeviceShareRepo(l.ctx).Insert(l.ctx, &po)
	if err != nil {
		return nil, err
	}
	return &dm.WithID{Id: po.ID}, nil
}