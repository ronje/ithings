package devicemsglogic

import (
	"context"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/shadow"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ShadowIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShadowIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShadowIndexLogic {
	return &ShadowIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取设备影子列表
func (l *ShadowIndexLogic) ShadowIndex(in *dm.PropertyLogLatestIndexReq) (*dm.ShadowIndexResp, error) {
	_, err := logic.SchemaAccess(l.ctx, l.svcCtx, def.AuthRead, devices.Core{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}, nil)
	if err != nil {
		return nil, err
	}
	sr := relationDB.NewShadowRepo(l.ctx)
	srs, err := sr.FindByFilter(l.ctx, shadow.Filter{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
		DataIDs:    in.DataIDs,
	})
	if err != nil {
		return nil, err
	}
	var index []*dm.ShadowIndex
	for _, v := range srs {
		index = append(index, &dm.ShadowIndex{
			DataID:            v.DataID,
			Value:             v.Value,
			UpdatedDeviceTime: utils.ToInt64(v.UpdatedDeviceTime),
		})
	}
	return &dm.ShadowIndexResp{List: index}, nil
}
