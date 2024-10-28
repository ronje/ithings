package dmExport

import (
	"context"

	"gitee.com/unitedrhino/share/caches"
	"gitee.com/unitedrhino/share/devices"
	"gitee.com/unitedrhino/share/domain/schema"
	"gitee.com/unitedrhino/share/eventBus"
	"gitee.com/unitedrhino/things/service/dmsvr/client/devicemanage"
	"gitee.com/unitedrhino/things/service/dmsvr/client/productmanage"
	"gitee.com/unitedrhino/things/service/dmsvr/client/userdevice"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/userShared"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
)

type ProductCacheT = *caches.Cache[dm.ProductInfo, string]

func NewProductInfoCache(pm productmanage.ProductManage, fastEvent *eventBus.FastEvent) (ProductCacheT, error) {
	return caches.NewCache(caches.CacheConfig[dm.ProductInfo, string]{
		KeyType:   eventBus.ServerCacheKeyDmProduct,
		FastEvent: fastEvent,
		GetData: func(ctx context.Context, key string) (*dm.ProductInfo, error) {
			ret, err := pm.ProductInfoRead(ctx, &dm.ProductInfoReadReq{ProductID: key, WithCategory: true, WithProtocol: true})
			return ret, err
		},
	})
}

type DeviceCacheT = *caches.Cache[dm.DeviceInfo, devices.Core]

func NewDeviceInfoCache(devM devicemanage.DeviceManage, fastEvent *eventBus.FastEvent) (DeviceCacheT, error) {
	return caches.NewCache(caches.CacheConfig[dm.DeviceInfo, devices.Core]{
		KeyType:   eventBus.ServerCacheKeyDmDevice,
		FastEvent: fastEvent,
		GetData: func(ctx context.Context, key devices.Core) (*dm.DeviceInfo, error) {
			ret, err := devM.DeviceInfoRead(ctx, &dm.DeviceInfoReadReq{ProductID: key.ProductID, DeviceName: key.DeviceName})
			return ret, err
		},
	})
}

//	type UserShareKey struct {
//		ProductID  string `json:"productID"`  //产品id
//		DeviceName string `json:"deviceName"` //设备名称
//		SharedUserID int64 `json:"sharedUserID"`
//	}
type UserShareCacheT = *caches.Cache[dm.UserDeviceShareInfo, userShared.UserShareKey]

func NewUserShareCache(devM userdevice.UserDevice, fastEvent *eventBus.FastEvent) (UserShareCacheT, error) {
	return caches.NewCache(caches.CacheConfig[dm.UserDeviceShareInfo, userShared.UserShareKey]{
		KeyType:   eventBus.ServerCacheKeyDmUserShareDevice,
		FastEvent: fastEvent,
		GetData: func(ctx context.Context, key userShared.UserShareKey) (*dm.UserDeviceShareInfo, error) {
			ret, err := devM.UserDeviceShareRead(ctx, &dm.UserDeviceShareReadReq{
				Device: &dm.DeviceCore{
					ProductID:  key.ProductID,
					DeviceName: key.DeviceName,
				},
			})
			return ret, err
		},
	})
}

type SchemaCacheT = *caches.Cache[schema.Model, devices.Core]

func NewSchemaInfoCache(pm productmanage.ProductManage, fastEvent *eventBus.FastEvent) (SchemaCacheT, error) {
	return caches.NewCache(caches.CacheConfig[schema.Model, devices.Core]{
		KeyType:   eventBus.ServerCacheKeyDmSchema,
		FastEvent: fastEvent,
		Fmt: func(ctx context.Context, key devices.Core, data *schema.Model) {
			data.ValidateWithFmt()
		},
		GetData: func(ctx context.Context, key devices.Core) (*schema.Model, error) {
			info, err := pm.ProductSchemaTslRead(ctx, &dm.ProductSchemaTslReadReq{ProductID: key.ProductID})
			if err != nil {
				return nil, err
			}
			return schema.ValidateWithFmt([]byte(info.Tsl))
		},
	})
}
