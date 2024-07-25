// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package productmanage

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                     = dm.ActionRespReq
	ActionSendReq                     = dm.ActionSendReq
	ActionSendResp                    = dm.ActionSendResp
	CommonSchemaCreateReq             = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq              = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp             = dm.CommonSchemaIndexResp
	CommonSchemaInfo                  = dm.CommonSchemaInfo
	CommonSchemaUpdateReq             = dm.CommonSchemaUpdateReq
	CompareInt64                      = dm.CompareInt64
	CompareString                     = dm.CompareString
	CustomTopic                       = dm.CustomTopic
	DeviceCore                        = dm.DeviceCore
	DeviceCountInfo                   = dm.DeviceCountInfo
	DeviceCountReq                    = dm.DeviceCountReq
	DeviceCountResp                   = dm.DeviceCountResp
	DeviceGatewayBindDevice           = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq             = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp            = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq       = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiSaveReq         = dm.DeviceGatewayMultiSaveReq
	DeviceGatewaySign                 = dm.DeviceGatewaySign
	DeviceInfo                        = dm.DeviceInfo
	DeviceInfoBindReq                 = dm.DeviceInfoBindReq
	DeviceInfoCanBindReq              = dm.DeviceInfoCanBindReq
	DeviceInfoCount                   = dm.DeviceInfoCount
	DeviceInfoCountReq                = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq               = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp               = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq          = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                 = dm.DeviceInfoReadReq
	DeviceModuleVersion               = dm.DeviceModuleVersion
	DeviceModuleVersionIndexReq       = dm.DeviceModuleVersionIndexReq
	DeviceModuleVersionIndexResp      = dm.DeviceModuleVersionIndexResp
	DeviceModuleVersionReadReq        = dm.DeviceModuleVersionReadReq
	DeviceOnlineMultiFix              = dm.DeviceOnlineMultiFix
	DeviceOnlineMultiFixReq           = dm.DeviceOnlineMultiFixReq
	DeviceProfile                     = dm.DeviceProfile
	DeviceProfileIndexReq             = dm.DeviceProfileIndexReq
	DeviceProfileIndexResp            = dm.DeviceProfileIndexResp
	DeviceProfileReadReq              = dm.DeviceProfileReadReq
	DeviceTransferReq                 = dm.DeviceTransferReq
	DeviceTypeCountReq                = dm.DeviceTypeCountReq
	DeviceTypeCountResp               = dm.DeviceTypeCountResp
	Empty                             = dm.Empty
	EventLogIndexReq                  = dm.EventLogIndexReq
	EventLogIndexResp                 = dm.EventLogIndexResp
	EventLogInfo                      = dm.EventLogInfo
	Firmware                          = dm.Firmware
	FirmwareFile                      = dm.FirmwareFile
	FirmwareInfo                      = dm.FirmwareInfo
	FirmwareInfoDeleteReq             = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp            = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq              = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp             = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq               = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp              = dm.FirmwareInfoReadResp
	FirmwareResp                      = dm.FirmwareResp
	GatewayCanBindIndexReq            = dm.GatewayCanBindIndexReq
	GatewayCanBindIndexResp           = dm.GatewayCanBindIndexResp
	GatewayGetFoundReq                = dm.GatewayGetFoundReq
	GatewayNotifyBindSendReq          = dm.GatewayNotifyBindSendReq
	GroupDeviceIndexReq               = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp              = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq         = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq           = dm.GroupDeviceMultiSaveReq
	GroupInfo                         = dm.GroupInfo
	GroupInfoCreateReq                = dm.GroupInfoCreateReq
	GroupInfoIndexReq                 = dm.GroupInfoIndexReq
	GroupInfoIndexResp                = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                = dm.GroupInfoUpdateReq
	HubLogIndexReq                    = dm.HubLogIndexReq
	HubLogIndexResp                   = dm.HubLogIndexResp
	HubLogInfo                        = dm.HubLogInfo
	IDPath                            = dm.IDPath
	IDPathWithUpdate                  = dm.IDPathWithUpdate
	OtaFirmwareDeviceCancelReq        = dm.OtaFirmwareDeviceCancelReq
	OtaFirmwareDeviceConfirmReq       = dm.OtaFirmwareDeviceConfirmReq
	OtaFirmwareDeviceIndexReq         = dm.OtaFirmwareDeviceIndexReq
	OtaFirmwareDeviceIndexResp        = dm.OtaFirmwareDeviceIndexResp
	OtaFirmwareDeviceInfo             = dm.OtaFirmwareDeviceInfo
	OtaFirmwareDeviceRetryReq         = dm.OtaFirmwareDeviceRetryReq
	OtaFirmwareFile                   = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq           = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp          = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo               = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp               = dm.OtaFirmwareFileResp
	OtaFirmwareInfo                   = dm.OtaFirmwareInfo
	OtaFirmwareInfoCreateReq          = dm.OtaFirmwareInfoCreateReq
	OtaFirmwareInfoIndexReq           = dm.OtaFirmwareInfoIndexReq
	OtaFirmwareInfoIndexResp          = dm.OtaFirmwareInfoIndexResp
	OtaFirmwareInfoUpdateReq          = dm.OtaFirmwareInfoUpdateReq
	OtaFirmwareJobIndexReq            = dm.OtaFirmwareJobIndexReq
	OtaFirmwareJobIndexResp           = dm.OtaFirmwareJobIndexResp
	OtaFirmwareJobInfo                = dm.OtaFirmwareJobInfo
	OtaJobByDeviceIndexReq            = dm.OtaJobByDeviceIndexReq
	OtaJobDynamicInfo                 = dm.OtaJobDynamicInfo
	OtaJobStaticInfo                  = dm.OtaJobStaticInfo
	OtaModuleInfo                     = dm.OtaModuleInfo
	OtaModuleInfoIndexReq             = dm.OtaModuleInfoIndexReq
	OtaModuleInfoIndexResp            = dm.OtaModuleInfoIndexResp
	OtaPromptIndexReq                 = dm.OtaPromptIndexReq
	OtaPromptIndexResp                = dm.OtaPromptIndexResp
	PageInfo                          = dm.PageInfo
	PageInfo_OrderBy                  = dm.PageInfo_OrderBy
	Point                             = dm.Point
	ProductCategory                   = dm.ProductCategory
	ProductCategoryIndexReq           = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp          = dm.ProductCategoryIndexResp
	ProductCategorySchemaIndexReq     = dm.ProductCategorySchemaIndexReq
	ProductCategorySchemaIndexResp    = dm.ProductCategorySchemaIndexResp
	ProductCategorySchemaMultiSaveReq = dm.ProductCategorySchemaMultiSaveReq
	ProductCustom                     = dm.ProductCustom
	ProductCustomReadReq              = dm.ProductCustomReadReq
	ProductInfo                       = dm.ProductInfo
	ProductInfoDeleteReq              = dm.ProductInfoDeleteReq
	ProductInfoIndexReq               = dm.ProductInfoIndexReq
	ProductInfoIndexResp              = dm.ProductInfoIndexResp
	ProductInfoReadReq                = dm.ProductInfoReadReq
	ProductInitReq                    = dm.ProductInitReq
	ProductRemoteConfig               = dm.ProductRemoteConfig
	ProductSchemaCreateReq            = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq            = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq             = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp            = dm.ProductSchemaIndexResp
	ProductSchemaInfo                 = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq       = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq         = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq           = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp          = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq            = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq       = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp      = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg            = dm.PropertyControlSendMsg
	PropertyControlSendReq            = dm.PropertyControlSendReq
	PropertyControlSendResp           = dm.PropertyControlSendResp
	PropertyGetReportSendReq          = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp         = dm.PropertyGetReportSendResp
	PropertyLogIndexReq               = dm.PropertyLogIndexReq
	PropertyLogIndexResp              = dm.PropertyLogIndexResp
	PropertyLogInfo                   = dm.PropertyLogInfo
	PropertyLogLatestIndexReq         = dm.PropertyLogLatestIndexReq
	ProtocolConfigField               = dm.ProtocolConfigField
	ProtocolConfigInfo                = dm.ProtocolConfigInfo
	ProtocolInfo                      = dm.ProtocolInfo
	ProtocolInfoIndexReq              = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp             = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq             = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq              = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp             = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq           = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp          = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq            = dm.RemoteConfigPushAllReq
	RespReadReq                       = dm.RespReadReq
	RootCheckReq                      = dm.RootCheckReq
	SdkLogIndexReq                    = dm.SdkLogIndexReq
	SdkLogIndexResp                   = dm.SdkLogIndexResp
	SdkLogInfo                        = dm.SdkLogInfo
	SendLogIndexReq                   = dm.SendLogIndexReq
	SendLogIndexResp                  = dm.SendLogIndexResp
	SendLogInfo                       = dm.SendLogInfo
	SendMsgReq                        = dm.SendMsgReq
	SendMsgResp                       = dm.SendMsgResp
	SendOption                        = dm.SendOption
	ShadowIndex                       = dm.ShadowIndex
	ShadowIndexResp                   = dm.ShadowIndexResp
	SharePerm                         = dm.SharePerm
	StatusLogIndexReq                 = dm.StatusLogIndexReq
	StatusLogIndexResp                = dm.StatusLogIndexResp
	StatusLogInfo                     = dm.StatusLogInfo
	TimeRange                         = dm.TimeRange
	UserDeviceCollectSave             = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq           = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp          = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo               = dm.UserDeviceShareInfo
	UserDeviceShareMultiDeleteReq     = dm.UserDeviceShareMultiDeleteReq
	UserDeviceShareReadReq            = dm.UserDeviceShareReadReq
	WithID                            = dm.WithID
	WithIDChildren                    = dm.WithIDChildren
	WithIDCode                        = dm.WithIDCode
	WithProfile                       = dm.WithProfile

	ProductManage interface {
		ProductInit(ctx context.Context, in *ProductInitReq, opts ...grpc.CallOption) (*Empty, error)
		// 新增产品
		ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error)
		// 更新产品
		ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error)
		// 删除产品
		ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取产品信息列表
		ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error)
		// 获取产品信息详情
		ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error)
		// 更新产品物模型
		ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Empty, error)
		// 新增产品
		ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Empty, error)
		// 批量新增物模型,只新增没有的,已有的不处理
		ProductSchemaMultiCreate(ctx context.Context, in *ProductSchemaMultiCreateReq, opts ...grpc.CallOption) (*Empty, error)
		// 删除产品
		ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取产品信息列表
		ProductSchemaIndex(ctx context.Context, in *ProductSchemaIndexReq, opts ...grpc.CallOption) (*ProductSchemaIndexResp, error)
		// 删除产品
		ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取产品信息列表
		ProductSchemaTslRead(ctx context.Context, in *ProductSchemaTslReadReq, opts ...grpc.CallOption) (*ProductSchemaTslReadResp, error)
		// 脚本管理
		ProductCustomRead(ctx context.Context, in *ProductCustomReadReq, opts ...grpc.CallOption) (*ProductCustom, error)
		ProductCustomUpdate(ctx context.Context, in *ProductCustom, opts ...grpc.CallOption) (*Empty, error)
		// 新增产品
		ProductCategoryCreate(ctx context.Context, in *ProductCategory, opts ...grpc.CallOption) (*WithID, error)
		// 更新产品
		ProductCategoryUpdate(ctx context.Context, in *ProductCategory, opts ...grpc.CallOption) (*Empty, error)
		// 删除产品
		ProductCategoryDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		// 获取产品信息列表
		ProductCategoryIndex(ctx context.Context, in *ProductCategoryIndexReq, opts ...grpc.CallOption) (*ProductCategoryIndexResp, error)
		// 获取产品信息详情
		ProductCategoryRead(ctx context.Context, in *WithIDChildren, opts ...grpc.CallOption) (*ProductCategory, error)
		// 获取产品品类下的物模型列表,绑定的物模型会自动添加到该产品品类及子分类的产品中,并不支持删除
		ProductCategorySchemaIndex(ctx context.Context, in *ProductCategorySchemaIndexReq, opts ...grpc.CallOption) (*ProductCategorySchemaIndexResp, error)
		ProductCategorySchemaMultiUpdate(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		ProductCategorySchemaMultiCreate(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		ProductCategorySchemaMultiDelete(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultProductManage struct {
		cli zrpc.Client
	}

	directProductManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.ProductManageServer
	}
)

func NewProductManage(cli zrpc.Client) ProductManage {
	return &defaultProductManage{
		cli: cli,
	}
}

func NewDirectProductManage(svcCtx *svc.ServiceContext, svr dm.ProductManageServer) ProductManage {
	return &directProductManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultProductManage) ProductInit(ctx context.Context, in *ProductInitReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInit(ctx, in, opts...)
}

func (d *directProductManage) ProductInit(ctx context.Context, in *ProductInitReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductInit(ctx, in)
}

// 新增产品
func (m *defaultProductManage) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoCreate(ctx, in, opts...)
}

// 新增产品
func (d *directProductManage) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductInfoCreate(ctx, in)
}

// 更新产品
func (m *defaultProductManage) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoUpdate(ctx, in, opts...)
}

// 更新产品
func (d *directProductManage) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductInfoUpdate(ctx, in)
}

// 删除产品
func (m *defaultProductManage) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoDelete(ctx, in, opts...)
}

// 删除产品
func (d *directProductManage) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductInfoDelete(ctx, in)
}

// 获取产品信息列表
func (m *defaultProductManage) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoIndex(ctx, in, opts...)
}

// 获取产品信息列表
func (d *directProductManage) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	return d.svr.ProductInfoIndex(ctx, in)
}

// 获取产品信息详情
func (m *defaultProductManage) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoRead(ctx, in, opts...)
}

// 获取产品信息详情
func (d *directProductManage) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	return d.svr.ProductInfoRead(ctx, in)
}

// 更新产品物模型
func (m *defaultProductManage) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaUpdate(ctx, in, opts...)
}

// 更新产品物模型
func (d *directProductManage) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductSchemaUpdate(ctx, in)
}

// 新增产品
func (m *defaultProductManage) ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaCreate(ctx, in, opts...)
}

// 新增产品
func (d *directProductManage) ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductSchemaCreate(ctx, in)
}

// 批量新增物模型,只新增没有的,已有的不处理
func (m *defaultProductManage) ProductSchemaMultiCreate(ctx context.Context, in *ProductSchemaMultiCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaMultiCreate(ctx, in, opts...)
}

// 批量新增物模型,只新增没有的,已有的不处理
func (d *directProductManage) ProductSchemaMultiCreate(ctx context.Context, in *ProductSchemaMultiCreateReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductSchemaMultiCreate(ctx, in)
}

// 删除产品
func (m *defaultProductManage) ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaDelete(ctx, in, opts...)
}

// 删除产品
func (d *directProductManage) ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductSchemaDelete(ctx, in)
}

// 获取产品信息列表
func (m *defaultProductManage) ProductSchemaIndex(ctx context.Context, in *ProductSchemaIndexReq, opts ...grpc.CallOption) (*ProductSchemaIndexResp, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaIndex(ctx, in, opts...)
}

// 获取产品信息列表
func (d *directProductManage) ProductSchemaIndex(ctx context.Context, in *ProductSchemaIndexReq, opts ...grpc.CallOption) (*ProductSchemaIndexResp, error) {
	return d.svr.ProductSchemaIndex(ctx, in)
}

// 删除产品
func (m *defaultProductManage) ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaTslImport(ctx, in, opts...)
}

// 删除产品
func (d *directProductManage) ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductSchemaTslImport(ctx, in)
}

// 获取产品信息列表
func (m *defaultProductManage) ProductSchemaTslRead(ctx context.Context, in *ProductSchemaTslReadReq, opts ...grpc.CallOption) (*ProductSchemaTslReadResp, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaTslRead(ctx, in, opts...)
}

// 获取产品信息列表
func (d *directProductManage) ProductSchemaTslRead(ctx context.Context, in *ProductSchemaTslReadReq, opts ...grpc.CallOption) (*ProductSchemaTslReadResp, error) {
	return d.svr.ProductSchemaTslRead(ctx, in)
}

// 脚本管理
func (m *defaultProductManage) ProductCustomRead(ctx context.Context, in *ProductCustomReadReq, opts ...grpc.CallOption) (*ProductCustom, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCustomRead(ctx, in, opts...)
}

// 脚本管理
func (d *directProductManage) ProductCustomRead(ctx context.Context, in *ProductCustomReadReq, opts ...grpc.CallOption) (*ProductCustom, error) {
	return d.svr.ProductCustomRead(ctx, in)
}

func (m *defaultProductManage) ProductCustomUpdate(ctx context.Context, in *ProductCustom, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCustomUpdate(ctx, in, opts...)
}

func (d *directProductManage) ProductCustomUpdate(ctx context.Context, in *ProductCustom, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductCustomUpdate(ctx, in)
}

// 新增产品
func (m *defaultProductManage) ProductCategoryCreate(ctx context.Context, in *ProductCategory, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategoryCreate(ctx, in, opts...)
}

// 新增产品
func (d *directProductManage) ProductCategoryCreate(ctx context.Context, in *ProductCategory, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.ProductCategoryCreate(ctx, in)
}

// 更新产品
func (m *defaultProductManage) ProductCategoryUpdate(ctx context.Context, in *ProductCategory, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategoryUpdate(ctx, in, opts...)
}

// 更新产品
func (d *directProductManage) ProductCategoryUpdate(ctx context.Context, in *ProductCategory, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductCategoryUpdate(ctx, in)
}

// 删除产品
func (m *defaultProductManage) ProductCategoryDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategoryDelete(ctx, in, opts...)
}

// 删除产品
func (d *directProductManage) ProductCategoryDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductCategoryDelete(ctx, in)
}

// 获取产品信息列表
func (m *defaultProductManage) ProductCategoryIndex(ctx context.Context, in *ProductCategoryIndexReq, opts ...grpc.CallOption) (*ProductCategoryIndexResp, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategoryIndex(ctx, in, opts...)
}

// 获取产品信息列表
func (d *directProductManage) ProductCategoryIndex(ctx context.Context, in *ProductCategoryIndexReq, opts ...grpc.CallOption) (*ProductCategoryIndexResp, error) {
	return d.svr.ProductCategoryIndex(ctx, in)
}

// 获取产品信息详情
func (m *defaultProductManage) ProductCategoryRead(ctx context.Context, in *WithIDChildren, opts ...grpc.CallOption) (*ProductCategory, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategoryRead(ctx, in, opts...)
}

// 获取产品信息详情
func (d *directProductManage) ProductCategoryRead(ctx context.Context, in *WithIDChildren, opts ...grpc.CallOption) (*ProductCategory, error) {
	return d.svr.ProductCategoryRead(ctx, in)
}

// 获取产品品类下的物模型列表,绑定的物模型会自动添加到该产品品类及子分类的产品中,并不支持删除
func (m *defaultProductManage) ProductCategorySchemaIndex(ctx context.Context, in *ProductCategorySchemaIndexReq, opts ...grpc.CallOption) (*ProductCategorySchemaIndexResp, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategorySchemaIndex(ctx, in, opts...)
}

// 获取产品品类下的物模型列表,绑定的物模型会自动添加到该产品品类及子分类的产品中,并不支持删除
func (d *directProductManage) ProductCategorySchemaIndex(ctx context.Context, in *ProductCategorySchemaIndexReq, opts ...grpc.CallOption) (*ProductCategorySchemaIndexResp, error) {
	return d.svr.ProductCategorySchemaIndex(ctx, in)
}

func (m *defaultProductManage) ProductCategorySchemaMultiUpdate(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategorySchemaMultiUpdate(ctx, in, opts...)
}

func (d *directProductManage) ProductCategorySchemaMultiUpdate(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductCategorySchemaMultiUpdate(ctx, in)
}

func (m *defaultProductManage) ProductCategorySchemaMultiCreate(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategorySchemaMultiCreate(ctx, in, opts...)
}

func (d *directProductManage) ProductCategorySchemaMultiCreate(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductCategorySchemaMultiCreate(ctx, in)
}

func (m *defaultProductManage) ProductCategorySchemaMultiDelete(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCategorySchemaMultiDelete(ctx, in, opts...)
}

func (d *directProductManage) ProductCategorySchemaMultiDelete(ctx context.Context, in *ProductCategorySchemaMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProductCategorySchemaMultiDelete(ctx, in)
}
