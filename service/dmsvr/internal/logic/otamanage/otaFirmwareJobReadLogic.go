package otamanagelogic

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
)

type OtaFirmwareJobReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	OjDB *relationDB.OtaJobRepo
}

func NewOtaFirmwareJobReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtaFirmwareJobReadLogic {
	return &OtaFirmwareJobReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		OjDB:   relationDB.NewOtaJobRepo(ctx),
	}
}

// //获取设备所在的升级包升级批次列表
func (l *OtaFirmwareJobReadLogic) OtaFirmwareJobRead(in *dm.WithID) (*dm.OtaFirmwareJobInfo, error) {
	if err := ctxs.IsRoot(l.ctx); err != nil {
		return nil, err
	}
	l.ctx = ctxs.WithRoot(l.ctx)
	otaJob, err := l.OjDB.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Errorf("%s.JobInfo.JobInfoRead failure err=%+v", utils.FuncName(), err)
		return nil, err
	}
	var otaJobInfo = dm.OtaFirmwareJobInfo{Dynamic: &dm.OtaJobDynamicInfo{}, Static: &dm.OtaJobStaticInfo{}}
	utils.CopyE(&otaJobInfo, &otaJob)
	return &otaJobInfo, nil
}
