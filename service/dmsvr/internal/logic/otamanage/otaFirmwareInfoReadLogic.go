package otamanagelogic

import (
	"context"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type OtaFirmwareInfoReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PiDB  *relationDB.ProductInfoRepo
	OfDB  *relationDB.OtaFirmwareInfoRepo
	OffDB *relationDB.OtaFirmwareFileRepo
}

func NewOtaFirmwareInfoReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtaFirmwareInfoReadLogic {
	return &OtaFirmwareInfoReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PiDB:   relationDB.NewProductInfoRepo(ctx),
		OfDB:   relationDB.NewOtaFirmwareInfoRepo(ctx),
		OffDB:  relationDB.NewOtaFirmwareFileRepo(ctx),
	}
}

// 查询升级包
func (l *OtaFirmwareInfoReadLogic) OtaFirmwareInfoRead(in *dm.WithID) (*dm.OtaFirmwareInfo, error) {
	otaFirmware, err := l.OfDB.FindOneByFilter(l.ctx, relationDB.OtaFirmwareInfoFilter{ID: in.Id, WithFiles: true})
	if err != nil {
		l.Errorf("%s.Query OTAFirmware err=%v", utils.FuncName(), err)
		return nil, err
	}
	return ToFirmwareInfoPb(l.ctx, l.svcCtx, otaFirmware), nil
}