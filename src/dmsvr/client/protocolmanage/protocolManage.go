// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package protocolmanage

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CustomTopic                        = dm.CustomTopic
	DeviceCore                         = dm.DeviceCore
	DeviceGatewayBindDevice            = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq              = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp             = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq        = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq        = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                  = dm.DeviceGatewaySign
	DeviceInfo                         = dm.DeviceInfo
	DeviceInfoCountReq                 = dm.DeviceInfoCountReq
	DeviceInfoCountResp                = dm.DeviceInfoCountResp
	DeviceInfoDeleteReq                = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                 = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                = dm.DeviceInfoIndexResp
	DeviceInfoReadReq                  = dm.DeviceInfoReadReq
	DeviceTypeCountReq                 = dm.DeviceTypeCountReq
	DeviceTypeCountResp                = dm.DeviceTypeCountResp
	DynamicUpgradeJobReq               = dm.DynamicUpgradeJobReq
	EventIndex                         = dm.EventIndex
	EventIndexResp                     = dm.EventIndexResp
	EventLogIndexReq                   = dm.EventLogIndexReq
	Firmware                           = dm.Firmware
	FirmwareFile                       = dm.FirmwareFile
	FirmwareInfo                       = dm.FirmwareInfo
	FirmwareInfoDeleteReq              = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp             = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq               = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp              = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq                = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp               = dm.FirmwareInfoReadResp
	FirmwareResp                       = dm.FirmwareResp
	GetPropertyReplyReq                = dm.GetPropertyReplyReq
	GetPropertyReplyResp               = dm.GetPropertyReplyResp
	GroupDeviceIndexReq                = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp               = dm.GroupDeviceIndexResp
	GroupDeviceMultiCreateReq          = dm.GroupDeviceMultiCreateReq
	GroupDeviceMultiDeleteReq          = dm.GroupDeviceMultiDeleteReq
	GroupInfo                          = dm.GroupInfo
	GroupInfoCreateReq                 = dm.GroupInfoCreateReq
	GroupInfoDeleteReq                 = dm.GroupInfoDeleteReq
	GroupInfoIndexReq                  = dm.GroupInfoIndexReq
	GroupInfoIndexResp                 = dm.GroupInfoIndexResp
	GroupInfoReadReq                   = dm.GroupInfoReadReq
	GroupInfoUpdateReq                 = dm.GroupInfoUpdateReq
	HubLogIndex                        = dm.HubLogIndex
	HubLogIndexReq                     = dm.HubLogIndexReq
	HubLogIndexResp                    = dm.HubLogIndexResp
	JobReq                             = dm.JobReq
	MultiSendPropertyReq               = dm.MultiSendPropertyReq
	MultiSendPropertyResp              = dm.MultiSendPropertyResp
	OTAModuleDeleteReq                 = dm.OTAModuleDeleteReq
	OTAModuleDetail                    = dm.OTAModuleDetail
	OTAModuleIndexReq                  = dm.OTAModuleIndexReq
	OTAModuleIndexResp                 = dm.OTAModuleIndexResp
	OTAModuleReq                       = dm.OTAModuleReq
	OTAModuleVersionsIndexResp         = dm.OTAModuleVersionsIndexResp
	OTATaskByDeviceCancelReq           = dm.OTATaskByDeviceCancelReq
	OTATaskByDeviceNameReq             = dm.OTATaskByDeviceNameReq
	OTATaskByJobCancelReq              = dm.OTATaskByJobCancelReq
	OTATaskByJobIndexReq               = dm.OTATaskByJobIndexReq
	OTATaskConfirmReq                  = dm.OTATaskConfirmReq
	OTATaskReUpgradeReq                = dm.OTATaskReUpgradeReq
	OTAUnfinishedTaskByDeviceIndexReq  = dm.OTAUnfinishedTaskByDeviceIndexReq
	OTAUnfinishedTaskByDeviceIndexResp = dm.OTAUnfinishedTaskByDeviceIndexResp
	OtaCommonResp                      = dm.OtaCommonResp
	OtaFirmwareCreateReq               = dm.OtaFirmwareCreateReq
	OtaFirmwareDeleteReq               = dm.OtaFirmwareDeleteReq
	OtaFirmwareDeviceInfoReq           = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp          = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile                    = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq            = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp           = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo                = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                 = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp                = dm.OtaFirmwareFileResp
	OtaFirmwareIndexReq                = dm.OtaFirmwareIndexReq
	OtaFirmwareIndexResp               = dm.OtaFirmwareIndexResp
	OtaFirmwareInfo                    = dm.OtaFirmwareInfo
	OtaFirmwareReadReq                 = dm.OtaFirmwareReadReq
	OtaFirmwareReadResp                = dm.OtaFirmwareReadResp
	OtaFirmwareResp                    = dm.OtaFirmwareResp
	OtaFirmwareUpdateReq               = dm.OtaFirmwareUpdateReq
	OtaFirmwareVerifyReq               = dm.OtaFirmwareVerifyReq
	OtaJobByDeviceIndexReq             = dm.OtaJobByDeviceIndexReq
	OtaJobByFirmwareIndexReq           = dm.OtaJobByFirmwareIndexReq
	OtaJobInfo                         = dm.OtaJobInfo
	OtaJobInfoIndexResp                = dm.OtaJobInfoIndexResp
	OtaModuleInfo                      = dm.OtaModuleInfo
	OtaPageInfo                        = dm.OtaPageInfo
	OtaPromptIndexReq                  = dm.OtaPromptIndexReq
	OtaPromptIndexResp                 = dm.OtaPromptIndexResp
	OtaTaskBatchReq                    = dm.OtaTaskBatchReq
	OtaTaskBatchResp                   = dm.OtaTaskBatchResp
	OtaTaskByJobIndexResp              = dm.OtaTaskByJobIndexResp
	OtaTaskCancleReq                   = dm.OtaTaskCancleReq
	OtaTaskCreatResp                   = dm.OtaTaskCreatResp
	OtaTaskCreateReq                   = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq             = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq              = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp             = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo                  = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq            = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq               = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq                    = dm.OtaTaskIndexReq
	OtaTaskIndexResp                   = dm.OtaTaskIndexResp
	OtaTaskInfo                        = dm.OtaTaskInfo
	OtaTaskReadReq                     = dm.OtaTaskReadReq
	OtaTaskReadResp                    = dm.OtaTaskReadResp
	OtaUpTaskInfo                      = dm.OtaUpTaskInfo
	PageInfo                           = dm.PageInfo
	PageInfo_OrderBy                   = dm.PageInfo_OrderBy
	Point                              = dm.Point
	ProductCustom                      = dm.ProductCustom
	ProductCustomReadReq               = dm.ProductCustomReadReq
	ProductInfo                        = dm.ProductInfo
	ProductInfoDeleteReq               = dm.ProductInfoDeleteReq
	ProductInfoIndexReq                = dm.ProductInfoIndexReq
	ProductInfoIndexResp               = dm.ProductInfoIndexResp
	ProductInfoReadReq                 = dm.ProductInfoReadReq
	ProductRemoteConfig                = dm.ProductRemoteConfig
	ProductSchemaCreateReq             = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq             = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq              = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp             = dm.ProductSchemaIndexResp
	ProductSchemaInfo                  = dm.ProductSchemaInfo
	ProductSchemaTslImportReq          = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq            = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp           = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq             = dm.ProductSchemaUpdateReq
	PropertyIndex                      = dm.PropertyIndex
	PropertyIndexResp                  = dm.PropertyIndexResp
	PropertyLatestIndexReq             = dm.PropertyLatestIndexReq
	PropertyLogIndexReq                = dm.PropertyLogIndexReq
	ProtocolInfo                       = dm.ProtocolInfo
	ProtocolInfoIndexReq               = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp              = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq              = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq               = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp              = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq            = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp           = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq             = dm.RemoteConfigPushAllReq
	RespActionReq                      = dm.RespActionReq
	RespReadReq                        = dm.RespReadReq
	Response                           = dm.Response
	RootCheckReq                       = dm.RootCheckReq
	SdkLogIndex                        = dm.SdkLogIndex
	SdkLogIndexReq                     = dm.SdkLogIndexReq
	SdkLogIndexResp                    = dm.SdkLogIndexResp
	SendActionReq                      = dm.SendActionReq
	SendActionResp                     = dm.SendActionResp
	SendMsgReq                         = dm.SendMsgReq
	SendMsgResp                        = dm.SendMsgResp
	SendOption                         = dm.SendOption
	SendPropertyMsg                    = dm.SendPropertyMsg
	SendPropertyReq                    = dm.SendPropertyReq
	SendPropertyResp                   = dm.SendPropertyResp
	ShadowIndex                        = dm.ShadowIndex
	ShadowIndexResp                    = dm.ShadowIndexResp
	StaticUpgradeDeviceInfo            = dm.StaticUpgradeDeviceInfo
	StaticUpgradeJobReq                = dm.StaticUpgradeJobReq
	Tag                                = dm.Tag
	UpgradeJobResp                     = dm.UpgradeJobResp
	VerifyOtaFirmwareReq               = dm.VerifyOtaFirmwareReq

	ProtocolManage interface {
		// 协议列表
		ProtocolInfoIndex(ctx context.Context, in *ProtocolInfoIndexReq, opts ...grpc.CallOption) (*ProtocolInfoIndexResp, error)
	}

	defaultProtocolManage struct {
		cli zrpc.Client
	}

	directProtocolManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.ProtocolManageServer
	}
)

func NewProtocolManage(cli zrpc.Client) ProtocolManage {
	return &defaultProtocolManage{
		cli: cli,
	}
}

func NewDirectProtocolManage(svcCtx *svc.ServiceContext, svr dm.ProtocolManageServer) ProtocolManage {
	return &directProtocolManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 协议列表
func (m *defaultProtocolManage) ProtocolInfoIndex(ctx context.Context, in *ProtocolInfoIndexReq, opts ...grpc.CallOption) (*ProtocolInfoIndexResp, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolInfoIndex(ctx, in, opts...)
}

// 协议列表
func (d *directProtocolManage) ProtocolInfoIndex(ctx context.Context, in *ProtocolInfoIndexReq, opts ...grpc.CallOption) (*ProtocolInfoIndexResp, error) {
	return d.svr.ProtocolInfoIndex(ctx, in)
}
