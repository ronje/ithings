// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package otafirmwaremanage

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                      = dm.ActionRespReq
	ActionSendReq                      = dm.ActionSendReq
	ActionSendResp                     = dm.ActionSendResp
	CommonSchemaCreateReq              = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq               = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp              = dm.CommonSchemaIndexResp
	CommonSchemaInfo                   = dm.CommonSchemaInfo
	CommonSchemaUpdateReq              = dm.CommonSchemaUpdateReq
	CustomTopic                        = dm.CustomTopic
	DeviceCore                         = dm.DeviceCore
	DeviceCountInfo                    = dm.DeviceCountInfo
	DeviceCountReq                     = dm.DeviceCountReq
	DeviceCountResp                    = dm.DeviceCountResp
	DeviceGatewayBindDevice            = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq              = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp             = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq        = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq        = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                  = dm.DeviceGatewaySign
	DeviceInfo                         = dm.DeviceInfo
	DeviceInfoCount                    = dm.DeviceInfoCount
	DeviceInfoCountReq                 = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq                = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                 = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq           = dm.DeviceInfoMultiUpdateReq
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
	GroupDeviceIndexReq                = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp               = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq          = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq            = dm.GroupDeviceMultiSaveReq
	GroupInfo                          = dm.GroupInfo
	GroupInfoCreateReq                 = dm.GroupInfoCreateReq
	GroupInfoIndexReq                  = dm.GroupInfoIndexReq
	GroupInfoIndexResp                 = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                 = dm.GroupInfoUpdateReq
	HubLogIndex                        = dm.HubLogIndex
	HubLogIndexReq                     = dm.HubLogIndexReq
	HubLogIndexResp                    = dm.HubLogIndexResp
	JobReq                             = dm.JobReq
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
	ProductCategory                    = dm.ProductCategory
	ProductCategoryIndexReq            = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp           = dm.ProductCategoryIndexResp
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
	PropertyControlMultiSendReq        = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp       = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg             = dm.PropertyControlSendMsg
	PropertyControlSendReq             = dm.PropertyControlSendReq
	PropertyControlSendResp            = dm.PropertyControlSendResp
	PropertyGetReportSendReq           = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp          = dm.PropertyGetReportSendResp
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
	RespReadReq                        = dm.RespReadReq
	Response                           = dm.Response
	RootCheckReq                       = dm.RootCheckReq
	SdkLogIndex                        = dm.SdkLogIndex
	SdkLogIndexReq                     = dm.SdkLogIndexReq
	SdkLogIndexResp                    = dm.SdkLogIndexResp
	SendMsgReq                         = dm.SendMsgReq
	SendMsgResp                        = dm.SendMsgResp
	SendOption                         = dm.SendOption
	ShadowIndex                        = dm.ShadowIndex
	ShadowIndexResp                    = dm.ShadowIndexResp
	StaticUpgradeDeviceInfo            = dm.StaticUpgradeDeviceInfo
	StaticUpgradeJobReq                = dm.StaticUpgradeJobReq
	Tag                                = dm.Tag
	TimeRange                          = dm.TimeRange
	UpgradeJobResp                     = dm.UpgradeJobResp
	VerifyOtaFirmwareReq               = dm.VerifyOtaFirmwareReq
	WithID                             = dm.WithID

	OTAFirmwareManage interface {
		// 添加升级包
		OtaFirmwareCreate(ctx context.Context, in *OtaFirmwareCreateReq, opts ...grpc.CallOption) (*OtaFirmwareResp, error)
		// 修改升级包
		OtaFirmwareUpdate(ctx context.Context, in *OtaFirmwareUpdateReq, opts ...grpc.CallOption) (*OtaFirmwareResp, error)
		// 删除升级包
		OtaFirmwareDelete(ctx context.Context, in *OtaFirmwareDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 升级包列表
		OtaFirmwareIndex(ctx context.Context, in *OtaFirmwareIndexReq, opts ...grpc.CallOption) (*OtaFirmwareIndexResp, error)
		// 查询升级包
		OtaFirmwareRead(ctx context.Context, in *OtaFirmwareReadReq, opts ...grpc.CallOption) (*OtaFirmwareReadResp, error)
	}

	defaultOTAFirmwareManage struct {
		cli zrpc.Client
	}

	directOTAFirmwareManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.OTAFirmwareManageServer
	}
)

func NewOTAFirmwareManage(cli zrpc.Client) OTAFirmwareManage {
	return &defaultOTAFirmwareManage{
		cli: cli,
	}
}

func NewDirectOTAFirmwareManage(svcCtx *svc.ServiceContext, svr dm.OTAFirmwareManageServer) OTAFirmwareManage {
	return &directOTAFirmwareManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 添加升级包
func (m *defaultOTAFirmwareManage) OtaFirmwareCreate(ctx context.Context, in *OtaFirmwareCreateReq, opts ...grpc.CallOption) (*OtaFirmwareResp, error) {
	client := dm.NewOTAFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareCreate(ctx, in, opts...)
}

// 添加升级包
func (d *directOTAFirmwareManage) OtaFirmwareCreate(ctx context.Context, in *OtaFirmwareCreateReq, opts ...grpc.CallOption) (*OtaFirmwareResp, error) {
	return d.svr.OtaFirmwareCreate(ctx, in)
}

// 修改升级包
func (m *defaultOTAFirmwareManage) OtaFirmwareUpdate(ctx context.Context, in *OtaFirmwareUpdateReq, opts ...grpc.CallOption) (*OtaFirmwareResp, error) {
	client := dm.NewOTAFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareUpdate(ctx, in, opts...)
}

// 修改升级包
func (d *directOTAFirmwareManage) OtaFirmwareUpdate(ctx context.Context, in *OtaFirmwareUpdateReq, opts ...grpc.CallOption) (*OtaFirmwareResp, error) {
	return d.svr.OtaFirmwareUpdate(ctx, in)
}

// 删除升级包
func (m *defaultOTAFirmwareManage) OtaFirmwareDelete(ctx context.Context, in *OtaFirmwareDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewOTAFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareDelete(ctx, in, opts...)
}

// 删除升级包
func (d *directOTAFirmwareManage) OtaFirmwareDelete(ctx context.Context, in *OtaFirmwareDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.OtaFirmwareDelete(ctx, in)
}

// 升级包列表
func (m *defaultOTAFirmwareManage) OtaFirmwareIndex(ctx context.Context, in *OtaFirmwareIndexReq, opts ...grpc.CallOption) (*OtaFirmwareIndexResp, error) {
	client := dm.NewOTAFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareIndex(ctx, in, opts...)
}

// 升级包列表
func (d *directOTAFirmwareManage) OtaFirmwareIndex(ctx context.Context, in *OtaFirmwareIndexReq, opts ...grpc.CallOption) (*OtaFirmwareIndexResp, error) {
	return d.svr.OtaFirmwareIndex(ctx, in)
}

// 查询升级包
func (m *defaultOTAFirmwareManage) OtaFirmwareRead(ctx context.Context, in *OtaFirmwareReadReq, opts ...grpc.CallOption) (*OtaFirmwareReadResp, error) {
	client := dm.NewOTAFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareRead(ctx, in, opts...)
}

// 查询升级包
func (d *directOTAFirmwareManage) OtaFirmwareRead(ctx context.Context, in *OtaFirmwareReadReq, opts ...grpc.CallOption) (*OtaFirmwareReadResp, error) {
	return d.svr.OtaFirmwareRead(ctx, in)
}
