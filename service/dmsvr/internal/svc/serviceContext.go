package svc

import (
	"context"
	"os"
	"time"

	"gitee.com/unitedrhino/core/service/syssvr/client/areamanage"
	"gitee.com/unitedrhino/core/service/syssvr/client/common"
	"gitee.com/unitedrhino/core/service/syssvr/client/datamanage"
	"gitee.com/unitedrhino/core/service/syssvr/client/projectmanage"
	"gitee.com/unitedrhino/core/service/syssvr/client/tenantmanage"
	"gitee.com/unitedrhino/core/service/syssvr/client/usermanage"
	"gitee.com/unitedrhino/core/service/syssvr/sysExport"
	"gitee.com/unitedrhino/core/service/timed/timedjobsvr/client/timedmanage"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/domain/deviceMsg/msgThing"
	ws "gitee.com/unitedrhino/share/websocket"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceLog"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/userShared"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/cache"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/event/publish/pubApp"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/event/publish/pubDev"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/relationDB"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/tdengine/schemaDataRepo"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/tdengine/sendLogRepo"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/tdengine/statusLogRepo"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"

	"gitee.com/unitedrhino/share/stores"

	"gitee.com/unitedrhino/share/caches"

	"gitee.com/unitedrhino/share/domain/schema"
	"gitee.com/unitedrhino/share/eventBus"
	"gitee.com/unitedrhino/share/oss"
	"gitee.com/unitedrhino/share/utils"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/config"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/tdengine/hubLogRepo"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/repo/tdengine/sdkLogRepo"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config

	PubDev pubDev.PubDev
	PubApp pubApp.PubApp

	OssClient            *oss.Client
	TimedM               timedmanage.TimedManage
	SchemaRepo           *caches.Cache[schema.Model, devices.Core]
	SchemaManaRepo       msgThing.SchemaDataRepo
	HubLogRepo           deviceLog.HubRepo
	StatusRepo           deviceLog.StatusRepo
	SendRepo             deviceLog.SendRepo
	SDKLogRepo           deviceLog.SDKRepo
	Cache                kv.Store
	DeviceStatus         *cache.DeviceStatus
	FastEvent            *eventBus.FastEvent
	AreaM                areamanage.AreaManage
	UserM                usermanage.UserManage
	DataM                datamanage.DataManage
	ProjectM             projectmanage.ProjectManage
	ProductCache         *caches.Cache[dm.ProductInfo, string]
	DeviceCache          *caches.Cache[dm.DeviceInfo, devices.Core]
	UserDeviceShare      *caches.Cache[dm.UserDeviceShareInfo, userShared.UserShareKey]
	UserMultiDeviceShare *caches.Cache[dm.UserDeviceShareMultiInfo, string]
	TenantCache          sysExport.TenantCacheT
	ProjectCache         sysExport.ProjectCacheT
	AreaCache            sysExport.AreaCacheT
	WebHook              *sysExport.Webhook
	Slot                 sysExport.SlotCacheT
	UserSubscribe        *ws.UserSubscribe
	GatewayCanBind       *cache.GatewayCanBind
	NodeID               int64
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		timedM   timedmanage.TimedManage
		areaM    areamanage.AreaManage
		projectM projectmanage.ProjectManage
	)
	stores.InitConn(c.Database)
	err := relationDB.Migrate(c.Database)
	if err != nil {
		logx.Error("dmsvr 初始化数据库错误 err", err)
		os.Exit(-1)
	}
	caches.InitStore(c.CacheRedis)
	nodeID := utils.GetNodeID(c.CacheRedis, c.Name)
	ca := kv.NewStore(c.CacheRedis)

	hubLogR := hubLogRepo.NewHubLogRepo(c.TSDB)
	sdkLogR := sdkLogRepo.NewSDKLogRepo(c.TSDB)
	statusR := statusLogRepo.NewStatusLogRepo(c.TSDB)
	sendR := sendLogRepo.NewSendLogRepo(c.TSDB)
	ossClient, err := oss.NewOssClient(c.OssConf)
	if err != nil {
		logx.Errorf("NewOss err err:%v", err)
		os.Exit(-1)
	}
	serverMsg, err := eventBus.NewFastEvent(c.Event, c.Name, nodeID)
	logx.Must(err)

	ccSchemaR, err := caches.NewCache(caches.CacheConfig[schema.Model, devices.Core]{
		KeyType:   eventBus.ServerCacheKeyDmSchema,
		FastEvent: serverMsg,
		GetData: func(ctx context.Context, key devices.Core) (*schema.Model, error) {
			db := relationDB.NewProductSchemaRepo(ctx)
			dbSchemas, err := db.FindByFilter(ctx, relationDB.ProductSchemaFilter{ProductID: key.ProductID}, nil)
			if err != nil {
				return nil, err
			}
			schemaModel := relationDB.ToSchemaDo(key.ProductID, dbSchemas)
			schemaModel.ValidateWithFmt()
			return schemaModel, nil
		},
		Fmt: func(ctx context.Context, key devices.Core, data *schema.Model) {
			data.ValidateWithFmt()
		},
		ExpireTime: 10 * time.Minute,
	})
	logx.Must(err)
	deviceDataR := schemaDataRepo.NewDeviceDataRepo(c.TSDB, ccSchemaR.GetData, ca)
	pd, err := pubDev.NewPubDev(serverMsg)
	if err != nil {
		logx.Error("NewPubDev err", err)
		os.Exit(-1)
	}
	pa, err := pubApp.NewPubApp(c.Event)
	if err != nil {
		logx.Error("NewPubApp err", err)
		os.Exit(-1)
	}

	timedM = timedmanage.NewTimedManage(zrpc.MustNewClient(c.TimedJobRpc.Conf))
	areaM = areamanage.NewAreaManage(zrpc.MustNewClient(c.SysRpc.Conf))
	userM := usermanage.NewUserManage(zrpc.MustNewClient(c.SysRpc.Conf))
	dataM := datamanage.NewDataManage(zrpc.MustNewClient(c.SysRpc.Conf))
	projectM = projectmanage.NewProjectManage(zrpc.MustNewClient(c.SysRpc.Conf))
	tenantCache, err := sysExport.NewTenantInfoCache(tenantmanage.NewTenantManage(zrpc.MustNewClient(c.SysRpc.Conf)), serverMsg)
	logx.Must(err)
	webHook, err := sysExport.NewTenantOpenWebhook(tenantmanage.NewTenantManage(zrpc.MustNewClient(c.SysRpc.Conf)), serverMsg)
	logx.Must(err)
	//webHook.Publish(ctxs.BindTenantCode(context.Background(), "default"),
	//	tenantOpenWebhook.CodeDmDeviceConn, application.ConnectMsg{Device: devices.Core{
	//		ProductID:  "123",
	//		DeviceName: "123",
	//	}, Timestamp: time.Now().UnixMilli()})
	Slot, err := sysExport.NewSlotCache(common.NewCommon(zrpc.MustNewClient(c.SysRpc.Conf)))
	logx.Must(err)
	projectC, err := sysExport.NewProjectInfoCache(projectmanage.NewProjectManage(zrpc.MustNewClient(c.SysRpc.Conf)), serverMsg)
	logx.Must(err)
	areaC, err := sysExport.NewAreaInfoCache(areamanage.NewAreaManage(zrpc.MustNewClient(c.SysRpc.Conf)), serverMsg)
	logx.Must(err)
	return &ServiceContext{
		FastEvent:      serverMsg,
		TenantCache:    tenantCache,
		Config:         c,
		OssClient:      ossClient,
		TimedM:         timedM,
		AreaM:          areaM,
		ProjectM:       projectM,
		PubApp:         pa,
		PubDev:         pd,
		Cache:          ca,
		UserM:          userM,
		DataM:          dataM,
		UserSubscribe:  ws.NewUserSubscribe(ca, serverMsg),
		SchemaRepo:     ccSchemaR,
		SchemaManaRepo: deviceDataR,
		DeviceStatus:   cache.NewDeviceStatus(ca),
		GatewayCanBind: cache.NewGatewayCanBind(ca),
		HubLogRepo:     hubLogR,
		SDKLogRepo:     sdkLogR,
		StatusRepo:     statusR,
		SendRepo:       sendR,
		WebHook:        webHook,
		NodeID:         nodeID,
		Slot:           Slot,
		ProjectCache:   projectC,
		AreaCache:      areaC,
	}
}

func (s *ServiceContext) WithDeviceTenant(ctx context.Context, dev devices.Core) context.Context {
	di, err := s.DeviceCache.GetData(ctx, dev)
	if err != nil {
		return ctx
	}
	return ctxs.BindTenantCode(ctx, di.TenantCode, di.ProjectID)
}
