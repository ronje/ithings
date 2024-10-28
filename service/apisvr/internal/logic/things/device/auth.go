package device

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/def"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"
	"gitee.com/unitedrhino/things/service/dgsvr/client/deviceauth"
	"gitee.com/unitedrhino/things/service/dgsvr/pb/dg"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"sync"
	"time"
)

var (
	protocolLink      = map[string]zrpc.Client{}
	protocolLinkMutex sync.RWMutex
)

func Init(svcCtx *svc.ServiceContext) {
	ctx := ctxs.WithRoot(context.Background())
	utils.Go(ctx, func() {
		ticket := time.NewTicker(time.Minute)
		run := func() {
			pi, err := svcCtx.ProtocolM.ProtocolInfoIndex(ctx, &dm.ProtocolInfoIndexReq{TransProtocol: def.ProtocolMqtt, NotCodes: []string{def.ProtocolCodeUnitedRhino}})
			if err != nil {
				logx.WithContext(ctx).Error(err)
				return
			}
			for _, pp := range pi.List {
				p := pp
				func() {
					protocolLinkMutex.Lock()
					defer protocolLinkMutex.Unlock()
					if _, ok := protocolLink[p.Code]; ok {
						return //已经连接上就不管了
					}
					var conf = svcCtx.Config.DgRpc.Conf
					if p.EtcdKey != "" {
						conf.Etcd = svcCtx.Config.Etcd
						conf.Etcd.Key = p.EtcdKey
					} else if p.Endpoints != nil {
						conf.Endpoints = p.Endpoints
					} else { //如果都没有配置,那么就不走这个服务校验
						return
					}
					cli, err := zrpc.NewClient(conf)
					if err == nil {
						val, ok := protocolLink[p.Code]
						protocolLink[p.Code] = cli
						if ok {
							time.Sleep(time.Second * 3) //避免刚取出来的连接失效,所以需要延时关闭
							val.(zrpc.Client).Conn().Close()
						}
					} else {
						val, ok := protocolLink[p.Code]
						if ok {
							time.Sleep(time.Second * 3) //避免刚取出来的连接失效,所以需要延时关闭
							val.(zrpc.Client).Conn().Close()
						}
					}
				}()
			}
		}
		run()
		for range ticket.C {
			run()
		}
	})
}

func ThirdProtoLoginAuth(ctx context.Context, svcCtx *svc.ServiceContext, req *types.DeviceAuthLoginReq, cert []byte) error {
	var wait sync.WaitGroup
	var succ bool
	var runCtx, cancel = context.WithCancel(ctx)
	protocolLinkMutex.RLock()
	defer protocolLinkMutex.RUnlock()
	for key, cli := range protocolLink {
		var start = time.Now()
		k := key
		c := cli
		wait.Add(1)
		utils.Go(ctx, func() {
			defer func() { wait.Done() }()
			da := deviceauth.NewDeviceAuth(c)
			_, err := da.LoginAuth(runCtx, &dg.LoginAuthReq{Username: req.Username, //用户名
				Password:    req.Password, //密码
				ClientID:    req.ClientID, //clientID
				Ip:          req.Ip,       //访问的ip地址
				Certificate: cert,         //客户端证书
			})
			if err == nil {
				succ = true
				logx.WithContext(runCtx).Infof("LoginAuth ProtocolKey:%v succ use:%v", k, time.Now().Sub(start))
				cancel()
			}
		})
	}
	wait.Wait()
	cancel()
	if succ {
		return nil
	}
	return errors.Permissions
}

func ThirdProtoAccessAuth(ctx context.Context, svcCtx *svc.ServiceContext, req *types.DeviceAuthAccessReq, action string) error {
	var wait sync.WaitGroup
	var succ bool
	var runCtx, cancel = context.WithCancel(ctx)
	protocolLinkMutex.RLock()
	defer protocolLinkMutex.RUnlock()
	for key, cli := range protocolLink {
		k := key
		c := cli
		wait.Add(1)
		utils.Go(ctx, func() {
			defer func() { wait.Done() }()
			da := deviceauth.NewDeviceAuth(c)
			_, err := da.AccessAuth(runCtx, &dg.AccessAuthReq{
				Username: req.Username,
				Topic:    req.Topic,
				ClientID: req.ClientID,
				Access:   action,
				Ip:       req.Ip,
			})
			if err == nil {
				logx.WithContext(runCtx).Infof("AccessAuth ProtocolKey:%v succ", k)
				succ = true
				cancel()
			}
		})
	}
	wait.Wait()
	cancel()
	if succ {
		return nil
	}
	return errors.Permissions
}
