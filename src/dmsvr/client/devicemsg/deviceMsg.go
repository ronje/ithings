// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package devicemsg

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CustomTopic                 = dm.CustomTopic
	DeviceCore                  = dm.DeviceCore
	DeviceGatewayBindDevice     = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq       = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp      = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign           = dm.DeviceGatewaySign
	DeviceInfo                  = dm.DeviceInfo
	DeviceInfoCountReq          = dm.DeviceInfoCountReq
	DeviceInfoCountResp         = dm.DeviceInfoCountResp
	DeviceInfoDeleteReq         = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq          = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp         = dm.DeviceInfoIndexResp
	DeviceInfoReadReq           = dm.DeviceInfoReadReq
	DeviceTypeCountReq          = dm.DeviceTypeCountReq
	DeviceTypeCountResp         = dm.DeviceTypeCountResp
	EventIndex                  = dm.EventIndex
	EventIndexResp              = dm.EventIndexResp
	EventLogIndexReq            = dm.EventLogIndexReq
	Firmware                    = dm.Firmware
	FirmwareInfo                = dm.FirmwareInfo
	FirmwareInfoDeleteReq       = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp      = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq        = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp       = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq         = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp        = dm.FirmwareInfoReadResp
	FirmwareResp                = dm.FirmwareResp
	GetPropertyReplyReq         = dm.GetPropertyReplyReq
	GetPropertyReplyResp        = dm.GetPropertyReplyResp
	GroupDeviceIndexReq         = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp        = dm.GroupDeviceIndexResp
	GroupDeviceMultiCreateReq   = dm.GroupDeviceMultiCreateReq
	GroupDeviceMultiDeleteReq   = dm.GroupDeviceMultiDeleteReq
	GroupInfo                   = dm.GroupInfo
	GroupInfoCreateReq          = dm.GroupInfoCreateReq
	GroupInfoDeleteReq          = dm.GroupInfoDeleteReq
	GroupInfoIndexReq           = dm.GroupInfoIndexReq
	GroupInfoIndexResp          = dm.GroupInfoIndexResp
	GroupInfoReadReq            = dm.GroupInfoReadReq
	GroupInfoUpdateReq          = dm.GroupInfoUpdateReq
	HubLogIndex                 = dm.HubLogIndex
	HubLogIndexReq              = dm.HubLogIndexReq
	HubLogIndexResp             = dm.HubLogIndexResp
	MultiSendPropertyReq        = dm.MultiSendPropertyReq
	MultiSendPropertyResp       = dm.MultiSendPropertyResp
	OtaCommonResp               = dm.OtaCommonResp
	OtaFirmwareDeviceInfoReq    = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp   = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile             = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq     = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp    = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo         = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq          = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp         = dm.OtaFirmwareFileResp
	OtaPageInfo                 = dm.OtaPageInfo
	OtaPromptIndexReq           = dm.OtaPromptIndexReq
	OtaPromptIndexResp          = dm.OtaPromptIndexResp
	OtaTaskBatchReq             = dm.OtaTaskBatchReq
	OtaTaskBatchResp            = dm.OtaTaskBatchResp
	OtaTaskCancleReq            = dm.OtaTaskCancleReq
	OtaTaskCreatResp            = dm.OtaTaskCreatResp
	OtaTaskCreateReq            = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq      = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq       = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp      = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo           = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq     = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq        = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq             = dm.OtaTaskIndexReq
	OtaTaskIndexResp            = dm.OtaTaskIndexResp
	OtaTaskInfo                 = dm.OtaTaskInfo
	OtaTaskReadReq              = dm.OtaTaskReadReq
	OtaTaskReadResp             = dm.OtaTaskReadResp
	PageInfo                    = dm.PageInfo
	PageInfo_OrderBy            = dm.PageInfo_OrderBy
	Point                       = dm.Point
	ProductCustom               = dm.ProductCustom
	ProductCustomReadReq        = dm.ProductCustomReadReq
	ProductInfo                 = dm.ProductInfo
	ProductInfoDeleteReq        = dm.ProductInfoDeleteReq
	ProductInfoIndexReq         = dm.ProductInfoIndexReq
	ProductInfoIndexResp        = dm.ProductInfoIndexResp
	ProductInfoReadReq          = dm.ProductInfoReadReq
	ProductRemoteConfig         = dm.ProductRemoteConfig
	ProductSchemaCreateReq      = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq      = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq       = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp      = dm.ProductSchemaIndexResp
	ProductSchemaInfo           = dm.ProductSchemaInfo
	ProductSchemaTslImportReq   = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq     = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp    = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq      = dm.ProductSchemaUpdateReq
	PropertyIndex               = dm.PropertyIndex
	PropertyIndexResp           = dm.PropertyIndexResp
	PropertyLatestIndexReq      = dm.PropertyLatestIndexReq
	PropertyLogIndexReq         = dm.PropertyLogIndexReq
	ProtocolInfo                = dm.ProtocolInfo
	ProtocolInfoIndexReq        = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp       = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq       = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq        = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp       = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq     = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp    = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq      = dm.RemoteConfigPushAllReq
	RespActionReq               = dm.RespActionReq
	RespReadReq                 = dm.RespReadReq
	Response                    = dm.Response
	RootCheckReq                = dm.RootCheckReq
	SdkLogIndex                 = dm.SdkLogIndex
	SdkLogIndexReq              = dm.SdkLogIndexReq
	SdkLogIndexResp             = dm.SdkLogIndexResp
	SendActionReq               = dm.SendActionReq
	SendActionResp              = dm.SendActionResp
	SendMsgReq                  = dm.SendMsgReq
	SendMsgResp                 = dm.SendMsgResp
	SendOption                  = dm.SendOption
	SendPropertyMsg             = dm.SendPropertyMsg
	SendPropertyReq             = dm.SendPropertyReq
	SendPropertyResp            = dm.SendPropertyResp
	ShadowIndex                 = dm.ShadowIndex
	ShadowIndexResp             = dm.ShadowIndexResp

	DeviceMsg interface {
		// 获取设备sdk调试日志
		SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error)
		// 获取设备调试信息记录登入登出,操作
		HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error)
		// 获取设备数据信息
		PropertyLatestIndex(ctx context.Context, in *PropertyLatestIndexReq, opts ...grpc.CallOption) (*PropertyIndexResp, error)
		// 获取设备数据信息
		PropertyLogIndex(ctx context.Context, in *PropertyLogIndexReq, opts ...grpc.CallOption) (*PropertyIndexResp, error)
		// 获取设备数据信息
		EventLogIndex(ctx context.Context, in *EventLogIndexReq, opts ...grpc.CallOption) (*EventIndexResp, error)
		// 获取设备影子列表
		ShadowIndex(ctx context.Context, in *PropertyLatestIndexReq, opts ...grpc.CallOption) (*ShadowIndexResp, error)
		// 主动触发单个设备ota升级推送
		OtaPromptIndex(ctx context.Context, in *OtaPromptIndexReq, opts ...grpc.CallOption) (*OtaPromptIndexResp, error)
	}

	defaultDeviceMsg struct {
		cli zrpc.Client
	}

	directDeviceMsg struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceMsgServer
	}
)

func NewDeviceMsg(cli zrpc.Client) DeviceMsg {
	return &defaultDeviceMsg{
		cli: cli,
	}
}

func NewDirectDeviceMsg(svcCtx *svc.ServiceContext, svr dm.DeviceMsgServer) DeviceMsg {
	return &directDeviceMsg{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 获取设备sdk调试日志
func (m *defaultDeviceMsg) SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.SdkLogIndex(ctx, in, opts...)
}

// 获取设备sdk调试日志
func (d *directDeviceMsg) SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error) {
	return d.svr.SdkLogIndex(ctx, in)
}

// 获取设备调试信息记录登入登出,操作
func (m *defaultDeviceMsg) HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.HubLogIndex(ctx, in, opts...)
}

// 获取设备调试信息记录登入登出,操作
func (d *directDeviceMsg) HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error) {
	return d.svr.HubLogIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) PropertyLatestIndex(ctx context.Context, in *PropertyLatestIndexReq, opts ...grpc.CallOption) (*PropertyIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.PropertyLatestIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) PropertyLatestIndex(ctx context.Context, in *PropertyLatestIndexReq, opts ...grpc.CallOption) (*PropertyIndexResp, error) {
	return d.svr.PropertyLatestIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) PropertyLogIndex(ctx context.Context, in *PropertyLogIndexReq, opts ...grpc.CallOption) (*PropertyIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.PropertyLogIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) PropertyLogIndex(ctx context.Context, in *PropertyLogIndexReq, opts ...grpc.CallOption) (*PropertyIndexResp, error) {
	return d.svr.PropertyLogIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) EventLogIndex(ctx context.Context, in *EventLogIndexReq, opts ...grpc.CallOption) (*EventIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.EventLogIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) EventLogIndex(ctx context.Context, in *EventLogIndexReq, opts ...grpc.CallOption) (*EventIndexResp, error) {
	return d.svr.EventLogIndex(ctx, in)
}

// 获取设备影子列表
func (m *defaultDeviceMsg) ShadowIndex(ctx context.Context, in *PropertyLatestIndexReq, opts ...grpc.CallOption) (*ShadowIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.ShadowIndex(ctx, in, opts...)
}

// 获取设备影子列表
func (d *directDeviceMsg) ShadowIndex(ctx context.Context, in *PropertyLatestIndexReq, opts ...grpc.CallOption) (*ShadowIndexResp, error) {
	return d.svr.ShadowIndex(ctx, in)
}

// 主动触发单个设备ota升级推送
func (m *defaultDeviceMsg) OtaPromptIndex(ctx context.Context, in *OtaPromptIndexReq, opts ...grpc.CallOption) (*OtaPromptIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.OtaPromptIndex(ctx, in, opts...)
}

// 主动触发单个设备ota升级推送
func (d *directDeviceMsg) OtaPromptIndex(ctx context.Context, in *OtaPromptIndexReq, opts ...grpc.CallOption) (*OtaPromptIndexResp, error) {
	return d.svr.OtaPromptIndex(ctx, in)
}