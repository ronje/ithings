package auth

import (
	"context"
	"encoding/base64"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/apisvr/internal/logic/things/device"
	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"
	"gitee.com/unitedrhino/things/service/dgsvr/pb/dg"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.DeviceAuthLoginReq) error {
	l.Infof("%s req=%+v", utils.FuncName(), req)
	var (
		cert []byte
		err  error
	)

	if req.Certificate != "" {
		cert, err = base64.StdEncoding.DecodeString(req.Certificate)
		if err != nil {
			return errors.Parameter.AddDetail("certificate can base64 decode")
		}

	}
	l.ctx = ctxs.WithRoot(l.ctx)
	_, err = l.svcCtx.DeviceM.RootCheck(l.ctx, &dm.RootCheckReq{
		Username:    req.Username,
		Password:    req.Password,
		ClientID:    req.ClientID,
		Ip:          req.Ip,
		Certificate: cert,
	})
	if err == nil { //root权限
		return nil
	}
	_, er := l.svcCtx.DeviceA.LoginAuth(l.ctx, &dg.LoginAuthReq{Username: req.Username, //用户名
		Password:    req.Password, //密码
		ClientID:    req.ClientID, //clientID
		Ip:          req.Ip,       //访问的ip地址
		Certificate: cert,         //客户端证书
	})
	if er == nil {
		return nil
	}
	err = device.ThirdProtoLoginAuth(l.ctx, l.svcCtx, req, cert)
	if err != nil {
		l.Errorf("authLogin req:%v iThings err:%v third err:%v", req, er, err)
	}
	return err
}
