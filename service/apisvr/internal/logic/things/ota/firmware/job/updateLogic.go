package job

import (
	"context"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.OtaFirmwareJobInfo) error {
	var firmwareUpdateReq dm.OtaFirmwareJobInfo
	_ = utils.CopyE(&firmwareUpdateReq, &req)
	logx.Infof("firmwareUpdateReq:%+v", &firmwareUpdateReq)
	_, err := l.svcCtx.OtaM.OtaFirmwareJobUpdate(l.ctx, &firmwareUpdateReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.OtaFirmwareUpdate req=%v err=%+v", utils.FuncName(), req, er)
		return er
	}
	return nil
}
