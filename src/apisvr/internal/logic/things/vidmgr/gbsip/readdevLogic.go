package gbsip

import (
	"context"
	"gitee.com/i-Things/core/shared/errors"
	"gitee.com/i-Things/core/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/vidsip/pb/sip"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReaddevLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReaddevLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReaddevLogic {
	return &ReaddevLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReaddevLogic) Readdev(req *types.VidmgrSipReadDevReq) (resp *types.VidmgrSipReadDevResp, err error) {
	// todo: add your logic here and delete this line
	vidResp, err := l.svcCtx.SipRpc.SipDeviceRead(l.ctx, &sip.SipDevReadReq{
		DeviceID: req.DeviceID,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s rpc.ManageVidmgr req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	apiResp := &types.VidmgrSipReadDevResp{
		CommonSipDevice: *VidmgrGbsipDeviceToApi(vidResp),
	}

	return apiResp, nil
}
