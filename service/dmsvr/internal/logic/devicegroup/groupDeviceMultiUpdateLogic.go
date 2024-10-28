package devicegrouplogic

import (
	"context"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDeviceMultiUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupDeviceMultiUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDeviceMultiUpdateLogic {
	return &GroupDeviceMultiUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新分组设备
func (l *GroupDeviceMultiUpdateLogic) GroupDeviceMultiUpdate(in *dm.GroupDeviceMultiSaveReq) (*dm.Empty, error) {
	t, err := relationDB.NewDeviceInfoRepo(l.ctx).CountByFilter(l.ctx, relationDB.DeviceFilter{Cores: logic.ToDeviceCores(in.List)})
	if err != nil {
		return nil, err
	}
	if int(t) != len(in.List) {
		return nil, errors.Duplicate.AddMsg("有被删除的设备请重试")
	}
	gi, err := relationDB.NewGroupInfoRepo(l.ctx).FindOne(l.ctx, in.GroupID)
	if err != nil {
		return nil, err
	}
	list := make([]*relationDB.DmGroupDevice, 0, len(in.List))
	for _, v := range in.List {
		list = append(list, &relationDB.DmGroupDevice{
			GroupID:    in.GroupID,
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
			AreaID:     gi.AreaID,
		})
	}
	err = relationDB.NewGroupDeviceRepo(l.ctx).MultiUpdate(l.ctx, in.GroupID, list)
	if err != nil {
		return nil, errors.Database.AddDetail(err)
	}
	relationDB.NewGroupInfoRepo(l.ctx).UpdateGroupDeviceCount(l.ctx, in.GroupID)
	return &dm.Empty{}, nil
}
