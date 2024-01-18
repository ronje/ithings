// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package rolemanage

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

	RoleManage interface {
		RoleInfoCreate(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*WithID, error)
		RoleInfoIndex(ctx context.Context, in *RoleInfoIndexReq, opts ...grpc.CallOption) (*RoleInfoIndexResp, error)
		RoleInfoUpdate(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*Response, error)
		RoleInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error)
		RoleMenuIndex(ctx context.Context, in *RoleMenuIndexReq, opts ...grpc.CallOption) (*RoleMenuIndexResp, error)
		RoleMenuMultiUpdate(ctx context.Context, in *RoleMenuMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleAppIndex(ctx context.Context, in *RoleAppIndexReq, opts ...grpc.CallOption) (*RoleAppIndexResp, error)
		RoleAppMultiUpdate(ctx context.Context, in *RoleAppMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleModuleIndex(ctx context.Context, in *RoleModuleIndexReq, opts ...grpc.CallOption) (*RoleModuleIndexResp, error)
		RoleModuleMultiUpdate(ctx context.Context, in *RoleModuleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleApiAuth(ctx context.Context, in *RoleApiAuthReq, opts ...grpc.CallOption) (*Response, error)
		RoleApiMultiUpdate(ctx context.Context, in *RoleApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		RoleApiIndex(ctx context.Context, in *RoleApiIndexReq, opts ...grpc.CallOption) (*RoleApiIndexResp, error)
	}

	defaultRoleManage struct {
		cli zrpc.Client
	}

	directRoleManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.RoleManageServer
	}
)

func NewRoleManage(cli zrpc.Client) RoleManage {
	return &defaultRoleManage{
		cli: cli,
	}
}

func NewDirectRoleManage(svcCtx *svc.ServiceContext, svr sys.RoleManageServer) RoleManage {
	return &directRoleManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultRoleManage) RoleInfoCreate(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleInfoCreate(ctx, in, opts...)
}

func (d *directRoleManage) RoleInfoCreate(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.RoleInfoCreate(ctx, in)
}

func (m *defaultRoleManage) RoleInfoIndex(ctx context.Context, in *RoleInfoIndexReq, opts ...grpc.CallOption) (*RoleInfoIndexResp, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleInfoIndex(ctx, in, opts...)
}

func (d *directRoleManage) RoleInfoIndex(ctx context.Context, in *RoleInfoIndexReq, opts ...grpc.CallOption) (*RoleInfoIndexResp, error) {
	return d.svr.RoleInfoIndex(ctx, in)
}

func (m *defaultRoleManage) RoleInfoUpdate(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleInfoUpdate(ctx, in, opts...)
}

func (d *directRoleManage) RoleInfoUpdate(ctx context.Context, in *RoleInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleInfoUpdate(ctx, in)
}

func (m *defaultRoleManage) RoleInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleInfoDelete(ctx, in, opts...)
}

func (d *directRoleManage) RoleInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleInfoDelete(ctx, in)
}

func (m *defaultRoleManage) RoleMenuIndex(ctx context.Context, in *RoleMenuIndexReq, opts ...grpc.CallOption) (*RoleMenuIndexResp, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleMenuIndex(ctx, in, opts...)
}

func (d *directRoleManage) RoleMenuIndex(ctx context.Context, in *RoleMenuIndexReq, opts ...grpc.CallOption) (*RoleMenuIndexResp, error) {
	return d.svr.RoleMenuIndex(ctx, in)
}

func (m *defaultRoleManage) RoleMenuMultiUpdate(ctx context.Context, in *RoleMenuMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleMenuMultiUpdate(ctx, in, opts...)
}

func (d *directRoleManage) RoleMenuMultiUpdate(ctx context.Context, in *RoleMenuMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleMenuMultiUpdate(ctx, in)
}

func (m *defaultRoleManage) RoleAppIndex(ctx context.Context, in *RoleAppIndexReq, opts ...grpc.CallOption) (*RoleAppIndexResp, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleAppIndex(ctx, in, opts...)
}

func (d *directRoleManage) RoleAppIndex(ctx context.Context, in *RoleAppIndexReq, opts ...grpc.CallOption) (*RoleAppIndexResp, error) {
	return d.svr.RoleAppIndex(ctx, in)
}

func (m *defaultRoleManage) RoleAppMultiUpdate(ctx context.Context, in *RoleAppMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleAppMultiUpdate(ctx, in, opts...)
}

func (d *directRoleManage) RoleAppMultiUpdate(ctx context.Context, in *RoleAppMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleAppMultiUpdate(ctx, in)
}

func (m *defaultRoleManage) RoleModuleIndex(ctx context.Context, in *RoleModuleIndexReq, opts ...grpc.CallOption) (*RoleModuleIndexResp, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleModuleIndex(ctx, in, opts...)
}

func (d *directRoleManage) RoleModuleIndex(ctx context.Context, in *RoleModuleIndexReq, opts ...grpc.CallOption) (*RoleModuleIndexResp, error) {
	return d.svr.RoleModuleIndex(ctx, in)
}

func (m *defaultRoleManage) RoleModuleMultiUpdate(ctx context.Context, in *RoleModuleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleModuleMultiUpdate(ctx, in, opts...)
}

func (d *directRoleManage) RoleModuleMultiUpdate(ctx context.Context, in *RoleModuleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleModuleMultiUpdate(ctx, in)
}

func (m *defaultRoleManage) RoleApiAuth(ctx context.Context, in *RoleApiAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleApiAuth(ctx, in, opts...)
}

func (d *directRoleManage) RoleApiAuth(ctx context.Context, in *RoleApiAuthReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleApiAuth(ctx, in)
}

func (m *defaultRoleManage) RoleApiMultiUpdate(ctx context.Context, in *RoleApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleApiMultiUpdate(ctx, in, opts...)
}

func (d *directRoleManage) RoleApiMultiUpdate(ctx context.Context, in *RoleApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.RoleApiMultiUpdate(ctx, in)
}

func (m *defaultRoleManage) RoleApiIndex(ctx context.Context, in *RoleApiIndexReq, opts ...grpc.CallOption) (*RoleApiIndexResp, error) {
	client := sys.NewRoleManageClient(m.cli.Conn())
	return client.RoleApiIndex(ctx, in, opts...)
}

func (d *directRoleManage) RoleApiIndex(ctx context.Context, in *RoleApiIndexReq, opts ...grpc.CallOption) (*RoleApiIndexResp, error) {
	return d.svr.RoleApiIndex(ctx, in)
}
