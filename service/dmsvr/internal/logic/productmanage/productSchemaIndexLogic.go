package productmanagelogic

import (
	"context"
	"gitee.com/unitedrhino/share/stores"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductSchemaIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PsDB *relationDB.ProductSchemaRepo
}

func NewProductSchemaIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductSchemaIndexLogic {
	return &ProductSchemaIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PsDB:   relationDB.NewProductSchemaRepo(ctx),
	}
}

// 获取产品信息列表
func (l *ProductSchemaIndexLogic) ProductSchemaIndex(in *dm.ProductSchemaIndexReq) (*dm.ProductSchemaIndexResp, error) {
	l.Infof("%s req=%v", utils.FuncName(), utils.Fmt(in))
	filter := relationDB.ProductSchemaFilter{
		ProductID:         in.ProductID,
		Type:              in.Type,
		Types:             in.Types,
		Tag:               in.Tag,
		Identifiers:       in.Identifiers,
		Name:              in.Name,
		IsCanSceneLinkage: in.IsCanSceneLinkage,
		FuncGroup:         in.FuncGroup,
		UserPerm:          in.UserPerm,
		PropertyMode:      in.PropertyMode,
	}
	schemas, err := l.PsDB.FindByFilter(l.ctx, filter, logic.ToPageInfo(in.Page).WithDefaultOrder(stores.OrderBy{
		Field: "order",
		Sort:  stores.OrderAsc,
	}))
	if err != nil {
		return nil, err
	}
	total, err := l.PsDB.CountByFilter(l.ctx, filter)
	if err != nil {
		return nil, err
	}
	list := make([]*dm.ProductSchemaInfo, 0, len(schemas))
	for _, s := range schemas {
		list = append(list, logic.ToProductSchemaRpc(s))
	}
	return &dm.ProductSchemaIndexResp{List: list, Total: total}, nil
}
