package rolemanagelogic

import (
	"context"
	appmanagelogic "github.com/i-Things/things/src/syssvr/internal/logic/appmanage"
	"github.com/i-Things/things/src/syssvr/internal/repo/relationDB"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAppIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAppIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAppIndexLogic {
	return &RoleAppIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAppIndexLogic) RoleAppIndex(in *sys.RoleAppIndexReq) (*sys.RoleAppIndexResp, error) {
	ra, err := relationDB.NewRoleAppRepo(l.ctx).FindByFilter(l.ctx, relationDB.RoleAppFilter{RoleID: in.Id}, nil)
	if err != nil {
		return nil, err
	}
	var appCodes []string
	for _, v := range ra {
		appCodes = append(appCodes, v.AppCode)
	}
	ais, err := relationDB.NewAppInfoRepo(l.ctx).FindByFilter(l.ctx, relationDB.AppInfoFilter{Codes: appCodes}, nil)
	if err != nil {
		return nil, err
	}

	return &sys.RoleAppIndexResp{List: appmanagelogic.ToAppInfosPb(ais), Total: int64(len(appCodes))}, nil
}