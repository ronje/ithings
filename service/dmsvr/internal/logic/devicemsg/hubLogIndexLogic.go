package devicemsglogic

import (
	"context"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceLog"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type HubLogIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHubLogIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HubLogIndexLogic {
	return &HubLogIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取设备调试信息记录登入登出,操作
func (l *HubLogIndexLogic) HubLogIndex(in *dm.HubLogIndexReq) (*dm.HubLogIndexResp, error) {
	_, err := logic.SchemaAccess(l.ctx, l.svcCtx, def.AuthRead, devices.Core{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}, nil)
	if err != nil {
		return nil, err
	}
	filter := deviceLog.HubFilter{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
		Actions:    in.Actions,
		Topics:     in.Topics,
		Content:    in.Content,
		RequestID:  in.RequestID,
	}
	page := def.PageInfo2{
		TimeStart: in.TimeStart,
		TimeEnd:   in.TimeEnd,
		Page:      in.Page.GetPage(),
		Size:      in.Page.GetSize(),
	}
	total, err := l.svcCtx.HubLogRepo.GetCountLog(l.ctx, filter, page)
	if err != nil {
		return nil, errors.Database.AddDetail(err)
	}
	logs, err := l.svcCtx.HubLogRepo.GetDeviceLog(l.ctx, filter, page)
	if err != nil {
		return nil, errors.Database.AddDetail(err)
	}

	var data []*dm.HubLogInfo
	for _, v := range logs {
		data = append(data, ToDataHubLogIndex(v))
	}
	return &dm.HubLogIndexResp{List: data, Total: total}, nil
}
