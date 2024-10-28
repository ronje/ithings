package collect

import (
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
)

func ToUserCollectDeviceSavePb(in *types.UserCollectDeviceSave) *dm.UserDeviceCollectSave {
	if in == nil {
		return nil
	}
	var devices []*dm.DeviceCore
	for _, v := range in.Devices {
		devices = append(devices, &dm.DeviceCore{
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
		})
	}
	return &dm.UserDeviceCollectSave{
		Devices: devices,
	}
}
