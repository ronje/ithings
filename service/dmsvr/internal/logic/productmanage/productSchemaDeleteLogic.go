package productmanagelogic

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/domain/schema"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductSchemaDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PsDB *relationDB.ProductSchemaRepo
}

func NewProductSchemaDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductSchemaDeleteLogic {
	return &ProductSchemaDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PsDB:   relationDB.NewProductSchemaRepo(ctx),
	}
}

// 删除产品
func (l *ProductSchemaDeleteLogic) ProductSchemaDelete(in *dm.ProductSchemaDeleteReq) (*dm.Empty, error) {
	if err := ctxs.IsRoot(l.ctx); err != nil {
		return nil, err
	}
	l.Infof("%s req=%v", utils.FuncName(), utils.Fmt(in))
	po, err := l.PsDB.FindOneByFilter(l.ctx, relationDB.ProductSchemaFilter{
		ProductID: in.ProductID, Identifiers: []string{in.Identifier},
	})
	if err != nil {
		if errors.Cmp(err, errors.NotFind) {
			return nil, errors.Parameter.AddMsg("标识符未找到")
		}
		return nil, err
	}
	if po.Tag == schema.TagRequired {
		return nil, errors.Parameter.AddMsg("必选物模型不能删除")
	}
	if schema.AffordanceType(po.Type) == schema.AffordanceTypeProperty {
		t, err := l.svcCtx.SchemaRepo.GetData(l.ctx, devices.Core{ProductID: in.ProductID})
		if err != nil {
			return nil, err
		}
		p, ok := t.Property[in.Identifier]
		if !ok {
			return nil, errors.Parameter.AddMsg("标识符未找到")
		}
		if err := l.svcCtx.SchemaManaRepo.DeleteProperty(l.ctx, p, in.ProductID, in.Identifier); err != nil {
			l.Errorf("%s.DeleteProperty failure,err:%v", utils.FuncName(), err)
			return nil, errors.Database.AddDetail(err)
		}
	}
	err = l.PsDB.DeleteByFilter(l.ctx, relationDB.ProductSchemaFilter{ID: po.ID})
	if err != nil {
		return nil, err
	}
	//清除缓存
	err = l.svcCtx.SchemaRepo.SetData(l.ctx, devices.Core{ProductID: in.ProductID}, nil)
	if err != nil {
		return nil, err
	}
	return &dm.Empty{}, nil
}
