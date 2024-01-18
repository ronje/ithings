// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package modulemanage

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiDeleteReq              = sys.ApiDeleteReq
	ApiInfo                   = sys.ApiInfo
	ApiInfoIndexReq           = sys.ApiInfoIndexReq
	ApiInfoIndexResp          = sys.ApiInfoIndexResp
	AppInfo                   = sys.AppInfo
	AppInfoIndexReq           = sys.AppInfoIndexReq
	AppInfoIndexResp          = sys.AppInfoIndexResp
	AppModuleIndexReq         = sys.AppModuleIndexReq
	AppModuleIndexResp        = sys.AppModuleIndexResp
	AppModuleMultiUpdateReq   = sys.AppModuleMultiUpdateReq
	AreaInfo                  = sys.AreaInfo
	AreaInfoIndexReq          = sys.AreaInfoIndexReq
	AreaInfoIndexResp         = sys.AreaInfoIndexResp
	AreaInfoReadReq           = sys.AreaInfoReadReq
	AreaWithID                = sys.AreaWithID
	AuthApiInfo               = sys.AuthApiInfo
	ConfigResp                = sys.ConfigResp
	DateRange                 = sys.DateRange
	JwtToken                  = sys.JwtToken
	LoginLogCreateReq         = sys.LoginLogCreateReq
	LoginLogIndexReq          = sys.LoginLogIndexReq
	LoginLogIndexResp         = sys.LoginLogIndexResp
	LoginLogInfo              = sys.LoginLogInfo
	Map                       = sys.Map
	MenuInfo                  = sys.MenuInfo
	MenuInfoIndexReq          = sys.MenuInfoIndexReq
	MenuInfoIndexResp         = sys.MenuInfoIndexResp
	ModuleInfo                = sys.ModuleInfo
	ModuleInfoIndexReq        = sys.ModuleInfoIndexReq
	ModuleInfoIndexResp       = sys.ModuleInfoIndexResp
	OperLogCreateReq          = sys.OperLogCreateReq
	OperLogIndexReq           = sys.OperLogIndexReq
	OperLogIndexResp          = sys.OperLogIndexResp
	OperLogInfo               = sys.OperLogInfo
	PageInfo                  = sys.PageInfo
	PageInfo_OrderBy          = sys.PageInfo_OrderBy
	Point                     = sys.Point
	ProjectInfo               = sys.ProjectInfo
	ProjectInfoIndexReq       = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp      = sys.ProjectInfoIndexResp
	ProjectWithID             = sys.ProjectWithID
	Response                  = sys.Response
	RoleApiAuthReq            = sys.RoleApiAuthReq
	RoleApiIndexReq           = sys.RoleApiIndexReq
	RoleApiIndexResp          = sys.RoleApiIndexResp
	RoleApiMultiUpdateReq     = sys.RoleApiMultiUpdateReq
	RoleAppIndexReq           = sys.RoleAppIndexReq
	RoleAppIndexResp          = sys.RoleAppIndexResp
	RoleAppMultiUpdateReq     = sys.RoleAppMultiUpdateReq
	RoleAppUpdateReq          = sys.RoleAppUpdateReq
	RoleInfo                  = sys.RoleInfo
	RoleInfoIndexReq          = sys.RoleInfoIndexReq
	RoleInfoIndexResp         = sys.RoleInfoIndexResp
	RoleMenuIndexReq          = sys.RoleMenuIndexReq
	RoleMenuIndexResp         = sys.RoleMenuIndexResp
	RoleMenuMultiUpdateReq    = sys.RoleMenuMultiUpdateReq
	RoleModuleIndexReq        = sys.RoleModuleIndexReq
	RoleModuleIndexResp       = sys.RoleModuleIndexResp
	RoleModuleMultiUpdateReq  = sys.RoleModuleMultiUpdateReq
	TenantApiInfo             = sys.TenantApiInfo
	TenantAppApiIndexReq      = sys.TenantAppApiIndexReq
	TenantAppApiIndexResp     = sys.TenantAppApiIndexResp
	TenantAppCreateReq        = sys.TenantAppCreateReq
	TenantAppIndexReq         = sys.TenantAppIndexReq
	TenantAppIndexResp        = sys.TenantAppIndexResp
	TenantAppMenu             = sys.TenantAppMenu
	TenantAppMenuIndexReq     = sys.TenantAppMenuIndexReq
	TenantAppMenuIndexResp    = sys.TenantAppMenuIndexResp
	TenantAppModule           = sys.TenantAppModule
	TenantAppMultiUpdateReq   = sys.TenantAppMultiUpdateReq
	TenantAppWithIDOrCode     = sys.TenantAppWithIDOrCode
	TenantInfo                = sys.TenantInfo
	TenantInfoCreateReq       = sys.TenantInfoCreateReq
	TenantInfoIndexReq        = sys.TenantInfoIndexReq
	TenantInfoIndexResp       = sys.TenantInfoIndexResp
	TenantModuleCreateReq     = sys.TenantModuleCreateReq
	TenantModuleIndexReq      = sys.TenantModuleIndexReq
	TenantModuleIndexResp     = sys.TenantModuleIndexResp
	TenantModuleWithIDOrCode  = sys.TenantModuleWithIDOrCode
	UserArea                  = sys.UserArea
	UserAreaApplyDealReq      = sys.UserAreaApplyDealReq
	UserAreaApplyIndexReq     = sys.UserAreaApplyIndexReq
	UserAreaApplyIndexResp    = sys.UserAreaApplyIndexResp
	UserAreaApplyInfo         = sys.UserAreaApplyInfo
	UserAreaApplyReq          = sys.UserAreaApplyReq
	UserAreaIndexReq          = sys.UserAreaIndexReq
	UserAreaIndexResp         = sys.UserAreaIndexResp
	UserAreaMultiDeleteReq    = sys.UserAreaMultiDeleteReq
	UserAreaMultiUpdateReq    = sys.UserAreaMultiUpdateReq
	UserCaptchaReq            = sys.UserCaptchaReq
	UserCaptchaResp           = sys.UserCaptchaResp
	UserChangePwdReq          = sys.UserChangePwdReq
	UserCheckTokenReq         = sys.UserCheckTokenReq
	UserCheckTokenResp        = sys.UserCheckTokenResp
	UserCreateResp            = sys.UserCreateResp
	UserForgetPwdReq          = sys.UserForgetPwdReq
	UserInfo                  = sys.UserInfo
	UserInfoCreateReq         = sys.UserInfoCreateReq
	UserInfoDeleteReq         = sys.UserInfoDeleteReq
	UserInfoIndexReq          = sys.UserInfoIndexReq
	UserInfoIndexResp         = sys.UserInfoIndexResp
	UserInfoReadReq           = sys.UserInfoReadReq
	UserLoginReq              = sys.UserLoginReq
	UserLoginResp             = sys.UserLoginResp
	UserProject               = sys.UserProject
	UserProjectIndexReq       = sys.UserProjectIndexReq
	UserProjectIndexResp      = sys.UserProjectIndexResp
	UserProjectMultiUpdateReq = sys.UserProjectMultiUpdateReq
	UserRegisterReq           = sys.UserRegisterReq
	UserRegisterResp          = sys.UserRegisterResp
	UserRoleIndexReq          = sys.UserRoleIndexReq
	UserRoleIndexResp         = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq    = sys.UserRoleMultiUpdateReq
	WithAppCodeID             = sys.WithAppCodeID
	WithID                    = sys.WithID
	WithIDCode                = sys.WithIDCode

	ModuleManage interface {
		ModuleInfoCreate(ctx context.Context, in *ModuleInfo, opts ...grpc.CallOption) (*WithID, error)
		ModuleInfoIndex(ctx context.Context, in *ModuleInfoIndexReq, opts ...grpc.CallOption) (*ModuleInfoIndexResp, error)
		ModuleInfoUpdate(ctx context.Context, in *ModuleInfo, opts ...grpc.CallOption) (*Response, error)
		ModuleInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error)
		ModuleInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*ModuleInfo, error)
		ModuleMenuCreate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*WithID, error)
		ModuleMenuIndex(ctx context.Context, in *MenuInfoIndexReq, opts ...grpc.CallOption) (*MenuInfoIndexResp, error)
		ModuleMenuUpdate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*Response, error)
		ModuleMenuDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error)
		ModuleApiCreate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*WithID, error)
		ModuleApiIndex(ctx context.Context, in *ApiInfoIndexReq, opts ...grpc.CallOption) (*ApiInfoIndexResp, error)
		ModuleApiUpdate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*Response, error)
		ModuleApiDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error)
	}

	defaultModuleManage struct {
		cli zrpc.Client
	}

	directModuleManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.ModuleManageServer
	}
)

func NewModuleManage(cli zrpc.Client) ModuleManage {
	return &defaultModuleManage{
		cli: cli,
	}
}

func NewDirectModuleManage(svcCtx *svc.ServiceContext, svr sys.ModuleManageServer) ModuleManage {
	return &directModuleManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultModuleManage) ModuleInfoCreate(ctx context.Context, in *ModuleInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleInfoCreate(ctx, in, opts...)
}

func (d *directModuleManage) ModuleInfoCreate(ctx context.Context, in *ModuleInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.ModuleInfoCreate(ctx, in)
}

func (m *defaultModuleManage) ModuleInfoIndex(ctx context.Context, in *ModuleInfoIndexReq, opts ...grpc.CallOption) (*ModuleInfoIndexResp, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleInfoIndex(ctx, in, opts...)
}

func (d *directModuleManage) ModuleInfoIndex(ctx context.Context, in *ModuleInfoIndexReq, opts ...grpc.CallOption) (*ModuleInfoIndexResp, error) {
	return d.svr.ModuleInfoIndex(ctx, in)
}

func (m *defaultModuleManage) ModuleInfoUpdate(ctx context.Context, in *ModuleInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleInfoUpdate(ctx, in, opts...)
}

func (d *directModuleManage) ModuleInfoUpdate(ctx context.Context, in *ModuleInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ModuleInfoUpdate(ctx, in)
}

func (m *defaultModuleManage) ModuleInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleInfoDelete(ctx, in, opts...)
}

func (d *directModuleManage) ModuleInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ModuleInfoDelete(ctx, in)
}

func (m *defaultModuleManage) ModuleInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*ModuleInfo, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleInfoRead(ctx, in, opts...)
}

func (d *directModuleManage) ModuleInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*ModuleInfo, error) {
	return d.svr.ModuleInfoRead(ctx, in)
}

func (m *defaultModuleManage) ModuleMenuCreate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleMenuCreate(ctx, in, opts...)
}

func (d *directModuleManage) ModuleMenuCreate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.ModuleMenuCreate(ctx, in)
}

func (m *defaultModuleManage) ModuleMenuIndex(ctx context.Context, in *MenuInfoIndexReq, opts ...grpc.CallOption) (*MenuInfoIndexResp, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleMenuIndex(ctx, in, opts...)
}

func (d *directModuleManage) ModuleMenuIndex(ctx context.Context, in *MenuInfoIndexReq, opts ...grpc.CallOption) (*MenuInfoIndexResp, error) {
	return d.svr.ModuleMenuIndex(ctx, in)
}

func (m *defaultModuleManage) ModuleMenuUpdate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleMenuUpdate(ctx, in, opts...)
}

func (d *directModuleManage) ModuleMenuUpdate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ModuleMenuUpdate(ctx, in)
}

func (m *defaultModuleManage) ModuleMenuDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleMenuDelete(ctx, in, opts...)
}

func (d *directModuleManage) ModuleMenuDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ModuleMenuDelete(ctx, in)
}

func (m *defaultModuleManage) ModuleApiCreate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleApiCreate(ctx, in, opts...)
}

func (d *directModuleManage) ModuleApiCreate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.ModuleApiCreate(ctx, in)
}

func (m *defaultModuleManage) ModuleApiIndex(ctx context.Context, in *ApiInfoIndexReq, opts ...grpc.CallOption) (*ApiInfoIndexResp, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleApiIndex(ctx, in, opts...)
}

func (d *directModuleManage) ModuleApiIndex(ctx context.Context, in *ApiInfoIndexReq, opts ...grpc.CallOption) (*ApiInfoIndexResp, error) {
	return d.svr.ModuleApiIndex(ctx, in)
}

func (m *defaultModuleManage) ModuleApiUpdate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleApiUpdate(ctx, in, opts...)
}

func (d *directModuleManage) ModuleApiUpdate(ctx context.Context, in *ApiInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ModuleApiUpdate(ctx, in)
}

func (m *defaultModuleManage) ModuleApiDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewModuleManageClient(m.cli.Conn())
	return client.ModuleApiDelete(ctx, in, opts...)
}

func (d *directModuleManage) ModuleApiDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ModuleApiDelete(ctx, in)
}
