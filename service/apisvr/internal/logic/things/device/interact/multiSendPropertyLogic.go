package interact

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"sync"

	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiSendPropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	retMsg []*types.DeviceInteractMultiSendPropertyMsg
	err    error
	mutex  sync.Mutex
}

func NewMultiSendPropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiSendPropertyLogic {
	return &MultiSendPropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctxs.WithDefaultRoot(ctx),
		svcCtx: svcCtx,
	}
}

func (l *MultiSendPropertyLogic) MultiSendProperty(req *types.DeviceInteractMultiSendPropertyReq) (resp *types.DeviceInteractMultiSendPropertyResp, err error) {
	return NewMultiSendPropertyControlLogic(l.ctx, l.svcCtx).MultiSendPropertyControl(req)
}
