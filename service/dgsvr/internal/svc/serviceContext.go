package svc

import (
	"context"
	"gitee.com/unitedrhino/core/service/timed/timedjobsvr/client/timedmanage"
	"gitee.com/unitedrhino/share/clients"
	"gitee.com/unitedrhino/share/conf"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/eventBus"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dgsvr/internal/config"
	"gitee.com/unitedrhino/things/service/dgsvr/internal/domain/custom"
	"gitee.com/unitedrhino/things/service/dgsvr/internal/repo/cache"
	"gitee.com/unitedrhino/things/service/dgsvr/internal/repo/event/publish/pubDev"
	"gitee.com/unitedrhino/things/service/dgsvr/internal/repo/event/publish/pubInner"
	"gitee.com/unitedrhino/things/service/dmsvr/client/devicemanage"
	"gitee.com/unitedrhino/things/service/dmsvr/client/productmanage"
	"gitee.com/unitedrhino/things/service/dmsvr/dmExport"
	"gitee.com/unitedrhino/things/service/dmsvr/dmdirect"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	PubDev       pubDev.PubDev
	MqttClient   *clients.MqttClient
	PubInner     pubInner.PubInner
	FastEvent    *eventBus.FastEvent
	TimedM       timedmanage.TimedManage
	ProductM     productmanage.ProductManage
	DeviceM      devicemanage.DeviceManage
	Script       custom.Repo
	ProductCache dmExport.ProductCacheT
	DeviceCache  dmExport.DeviceCacheT
	NodeID       int64
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		productM productmanage.ProductManage
		deviceM  devicemanage.DeviceManage
		timedM   timedmanage.TimedManage
	)
	/*
		// move to startup.PostInit()
			dl, err := pubDev.NewPubDev(c.DevLink)
			if err != nil {
				logx.Error("NewDevClient err", err)
				os.Exit(-1)
			}

			il, err := pubInner.NewPubInner(c.Event)
			if err != nil {
				logx.Error("NewInnerDevPub err", err)
				os.Exit(-1)
			}
	*/
	if c.DmRpc.Mode == conf.ClientModeGrpc {
		productM = productmanage.NewProductManage(zrpc.MustNewClient(c.DmRpc.Conf))
		deviceM = devicemanage.NewDeviceManage(zrpc.MustNewClient(c.DmRpc.Conf))
	} else {
		productM = dmdirect.NewProductManage(c.DmRpc.RunProxy)
		deviceM = dmdirect.NewDeviceManage(c.DmRpc.RunProxy)
	}
	scriptCache := cache.NewScriptRepo(func(ctx context.Context, productID string) (info *custom.Info, err error) {
		ret, err := productM.ProductCustomRead(ctx, &dm.ProductCustomReadReq{ProductID: productID})
		if err != nil {
			if errors.Cmp(err, errors.NotFind) { //如果是没找到
				return nil, nil
			}
			return nil, err
		}
		return &custom.Info{
			ProductID:       ret.ProductID,
			TransformScript: utils.ToNullString(ret.TransformScript),
			ScriptLang:      ret.ScriptLang,
		}, nil
	})
	nodeID := utils.GetNodeID(c.CacheRedis, c.Name)
	serverMsg, err := eventBus.NewFastEvent(c.Event, c.Name, nodeID)
	logx.Must(err)
	pc, err := dmExport.NewProductInfoCache(productM, serverMsg)
	logx.Must(err)
	dc, err := dmExport.NewDeviceInfoCache(deviceM, serverMsg)
	logx.Must(err)
	timedM = timedmanage.NewTimedManage(zrpc.MustNewClient(c.TimedJobRpc.Conf))
	return &ServiceContext{
		Config: c,
		// PubDev:   dl,
		// PubInner: il,
		FastEvent:    serverMsg,
		ProductM:     productM,
		DeviceM:      deviceM,
		Script:       scriptCache,
		ProductCache: pc,
		DeviceCache:  dc,
		NodeID:       nodeID,
		TimedM:       timedM,
	}
}
