// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package usermanage

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

	UserManage interface {
		UserInfoCreate(ctx context.Context, in *UserInfoCreateReq, opts ...grpc.CallOption) (*UserCreateResp, error)
		UserInfoIndex(ctx context.Context, in *UserInfoIndexReq, opts ...grpc.CallOption) (*UserInfoIndexResp, error)
		UserInfoUpdate(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Response, error)
		UserInfoRead(ctx context.Context, in *UserInfoReadReq, opts ...grpc.CallOption) (*UserInfo, error)
		UserInfoDelete(ctx context.Context, in *UserInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
		UserForgetPwd(ctx context.Context, in *UserForgetPwdReq, opts ...grpc.CallOption) (*Response, error)
		UserCaptcha(ctx context.Context, in *UserCaptchaReq, opts ...grpc.CallOption) (*UserCaptchaResp, error)
		UserCheckToken(ctx context.Context, in *UserCheckTokenReq, opts ...grpc.CallOption) (*UserCheckTokenResp, error)
		UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
		UserChangePwd(ctx context.Context, in *UserChangePwdReq, opts ...grpc.CallOption) (*Response, error)
		UserRoleIndex(ctx context.Context, in *UserRoleIndexReq, opts ...grpc.CallOption) (*UserRoleIndexResp, error)
		UserRoleMultiUpdate(ctx context.Context, in *UserRoleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		UserProjectMultiUpdate(ctx context.Context, in *UserProjectMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		UserProjectIndex(ctx context.Context, in *UserProjectIndexReq, opts ...grpc.CallOption) (*UserProjectIndexResp, error)
		UserAreaMultiUpdate(ctx context.Context, in *UserAreaMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		UserAreaIndex(ctx context.Context, in *UserAreaIndexReq, opts ...grpc.CallOption) (*UserAreaIndexResp, error)
		UserAreaMultiDelete(ctx context.Context, in *UserAreaMultiDeleteReq, opts ...grpc.CallOption) (*Response, error)
		UserAreaApply(ctx context.Context, in *UserAreaApplyReq, opts ...grpc.CallOption) (*Response, error)
		UserAreaApplyIndex(ctx context.Context, in *UserAreaApplyIndexReq, opts ...grpc.CallOption) (*UserAreaApplyIndexResp, error)
		UserAreaApplyDeal(ctx context.Context, in *UserAreaApplyDealReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUserManage struct {
		cli zrpc.Client
	}

	directUserManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.UserManageServer
	}
)

func NewUserManage(cli zrpc.Client) UserManage {
	return &defaultUserManage{
		cli: cli,
	}
}

func NewDirectUserManage(svcCtx *svc.ServiceContext, svr sys.UserManageServer) UserManage {
	return &directUserManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultUserManage) UserInfoCreate(ctx context.Context, in *UserInfoCreateReq, opts ...grpc.CallOption) (*UserCreateResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserInfoCreate(ctx, in, opts...)
}

func (d *directUserManage) UserInfoCreate(ctx context.Context, in *UserInfoCreateReq, opts ...grpc.CallOption) (*UserCreateResp, error) {
	return d.svr.UserInfoCreate(ctx, in)
}

func (m *defaultUserManage) UserInfoIndex(ctx context.Context, in *UserInfoIndexReq, opts ...grpc.CallOption) (*UserInfoIndexResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserInfoIndex(ctx, in, opts...)
}

func (d *directUserManage) UserInfoIndex(ctx context.Context, in *UserInfoIndexReq, opts ...grpc.CallOption) (*UserInfoIndexResp, error) {
	return d.svr.UserInfoIndex(ctx, in)
}

func (m *defaultUserManage) UserInfoUpdate(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserInfoUpdate(ctx, in, opts...)
}

func (d *directUserManage) UserInfoUpdate(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserInfoUpdate(ctx, in)
}

func (m *defaultUserManage) UserInfoRead(ctx context.Context, in *UserInfoReadReq, opts ...grpc.CallOption) (*UserInfo, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserInfoRead(ctx, in, opts...)
}

func (d *directUserManage) UserInfoRead(ctx context.Context, in *UserInfoReadReq, opts ...grpc.CallOption) (*UserInfo, error) {
	return d.svr.UserInfoRead(ctx, in)
}

func (m *defaultUserManage) UserInfoDelete(ctx context.Context, in *UserInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserInfoDelete(ctx, in, opts...)
}

func (d *directUserManage) UserInfoDelete(ctx context.Context, in *UserInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserInfoDelete(ctx, in)
}

func (m *defaultUserManage) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

func (d *directUserManage) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	return d.svr.UserLogin(ctx, in)
}

func (m *defaultUserManage) UserForgetPwd(ctx context.Context, in *UserForgetPwdReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserForgetPwd(ctx, in, opts...)
}

func (d *directUserManage) UserForgetPwd(ctx context.Context, in *UserForgetPwdReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserForgetPwd(ctx, in)
}

func (m *defaultUserManage) UserCaptcha(ctx context.Context, in *UserCaptchaReq, opts ...grpc.CallOption) (*UserCaptchaResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserCaptcha(ctx, in, opts...)
}

func (d *directUserManage) UserCaptcha(ctx context.Context, in *UserCaptchaReq, opts ...grpc.CallOption) (*UserCaptchaResp, error) {
	return d.svr.UserCaptcha(ctx, in)
}

func (m *defaultUserManage) UserCheckToken(ctx context.Context, in *UserCheckTokenReq, opts ...grpc.CallOption) (*UserCheckTokenResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserCheckToken(ctx, in, opts...)
}

func (d *directUserManage) UserCheckToken(ctx context.Context, in *UserCheckTokenReq, opts ...grpc.CallOption) (*UserCheckTokenResp, error) {
	return d.svr.UserCheckToken(ctx, in)
}

func (m *defaultUserManage) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserRegister(ctx, in, opts...)
}

func (d *directUserManage) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	return d.svr.UserRegister(ctx, in)
}

func (m *defaultUserManage) UserChangePwd(ctx context.Context, in *UserChangePwdReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserChangePwd(ctx, in, opts...)
}

func (d *directUserManage) UserChangePwd(ctx context.Context, in *UserChangePwdReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserChangePwd(ctx, in)
}

func (m *defaultUserManage) UserRoleIndex(ctx context.Context, in *UserRoleIndexReq, opts ...grpc.CallOption) (*UserRoleIndexResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserRoleIndex(ctx, in, opts...)
}

func (d *directUserManage) UserRoleIndex(ctx context.Context, in *UserRoleIndexReq, opts ...grpc.CallOption) (*UserRoleIndexResp, error) {
	return d.svr.UserRoleIndex(ctx, in)
}

func (m *defaultUserManage) UserRoleMultiUpdate(ctx context.Context, in *UserRoleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserRoleMultiUpdate(ctx, in, opts...)
}

func (d *directUserManage) UserRoleMultiUpdate(ctx context.Context, in *UserRoleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserRoleMultiUpdate(ctx, in)
}

func (m *defaultUserManage) UserProjectMultiUpdate(ctx context.Context, in *UserProjectMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserProjectMultiUpdate(ctx, in, opts...)
}

func (d *directUserManage) UserProjectMultiUpdate(ctx context.Context, in *UserProjectMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserProjectMultiUpdate(ctx, in)
}

func (m *defaultUserManage) UserProjectIndex(ctx context.Context, in *UserProjectIndexReq, opts ...grpc.CallOption) (*UserProjectIndexResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserProjectIndex(ctx, in, opts...)
}

func (d *directUserManage) UserProjectIndex(ctx context.Context, in *UserProjectIndexReq, opts ...grpc.CallOption) (*UserProjectIndexResp, error) {
	return d.svr.UserProjectIndex(ctx, in)
}

func (m *defaultUserManage) UserAreaMultiUpdate(ctx context.Context, in *UserAreaMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserAreaMultiUpdate(ctx, in, opts...)
}

func (d *directUserManage) UserAreaMultiUpdate(ctx context.Context, in *UserAreaMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserAreaMultiUpdate(ctx, in)
}

func (m *defaultUserManage) UserAreaIndex(ctx context.Context, in *UserAreaIndexReq, opts ...grpc.CallOption) (*UserAreaIndexResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserAreaIndex(ctx, in, opts...)
}

func (d *directUserManage) UserAreaIndex(ctx context.Context, in *UserAreaIndexReq, opts ...grpc.CallOption) (*UserAreaIndexResp, error) {
	return d.svr.UserAreaIndex(ctx, in)
}

func (m *defaultUserManage) UserAreaMultiDelete(ctx context.Context, in *UserAreaMultiDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserAreaMultiDelete(ctx, in, opts...)
}

func (d *directUserManage) UserAreaMultiDelete(ctx context.Context, in *UserAreaMultiDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserAreaMultiDelete(ctx, in)
}

func (m *defaultUserManage) UserAreaApply(ctx context.Context, in *UserAreaApplyReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserAreaApply(ctx, in, opts...)
}

func (d *directUserManage) UserAreaApply(ctx context.Context, in *UserAreaApplyReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserAreaApply(ctx, in)
}

func (m *defaultUserManage) UserAreaApplyIndex(ctx context.Context, in *UserAreaApplyIndexReq, opts ...grpc.CallOption) (*UserAreaApplyIndexResp, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserAreaApplyIndex(ctx, in, opts...)
}

func (d *directUserManage) UserAreaApplyIndex(ctx context.Context, in *UserAreaApplyIndexReq, opts ...grpc.CallOption) (*UserAreaApplyIndexResp, error) {
	return d.svr.UserAreaApplyIndex(ctx, in)
}

func (m *defaultUserManage) UserAreaApplyDeal(ctx context.Context, in *UserAreaApplyDealReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewUserManageClient(m.cli.Conn())
	return client.UserAreaApplyDeal(ctx, in, opts...)
}

func (d *directUserManage) UserAreaApplyDeal(ctx context.Context, in *UserAreaApplyDealReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.UserAreaApplyDeal(ctx, in)
}
