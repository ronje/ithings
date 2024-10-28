package otamanagelogic

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type OtaModuleInfoIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOtaModuleInfoIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtaModuleInfoIndexLogic {
	return &OtaModuleInfoIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OtaModuleInfoIndexLogic) OtaModuleInfoIndex(in *dm.OtaModuleInfoIndexReq) (*dm.OtaModuleInfoIndexResp, error) {
	if err := ctxs.IsRoot(l.ctx); err != nil {
		return nil, err
	}
	l.ctx = ctxs.WithRoot(l.ctx)
	f := relationDB.OtaModuleInfoFilter{ProductID: in.ProductID, Name: in.Name}
	total, err := relationDB.NewOtaModuleInfoRepo(l.ctx).CountByFilter(l.ctx, f)
	if err != nil {
		return nil, err
	}
	pos, err := relationDB.NewOtaModuleInfoRepo(l.ctx).FindByFilter(l.ctx, f, logic.ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	var list []*dm.OtaModuleInfo
	for _, v := range pos {
		list = append(list, utils.Copy[dm.OtaModuleInfo](v))
	}
	return &dm.OtaModuleInfoIndexResp{List: list, Total: total}, nil
}
