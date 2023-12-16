// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package common

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiDeleteReq                  = sys.ApiDeleteReq
	ApiInfo                       = sys.ApiInfo
	ApiInfoIndexReq               = sys.ApiInfoIndexReq
	ApiInfoIndexResp              = sys.ApiInfoIndexResp
	AppInfo                       = sys.AppInfo
	AppInfoIndexReq               = sys.AppInfoIndexReq
	AppInfoIndexResp              = sys.AppInfoIndexResp
	AreaInfo                      = sys.AreaInfo
	AreaInfoDeleteReq             = sys.AreaInfoDeleteReq
	AreaInfoIndexReq              = sys.AreaInfoIndexReq
	AreaInfoIndexResp             = sys.AreaInfoIndexResp
	AreaInfoReadReq               = sys.AreaInfoReadReq
	AreaInfoTreeReq               = sys.AreaInfoTreeReq
	AreaInfoTreeResp              = sys.AreaInfoTreeResp
	AuthApiInfo                   = sys.AuthApiInfo
	ConfigResp                    = sys.ConfigResp
	DateRange                     = sys.DateRange
	JwtToken                      = sys.JwtToken
	LoginLogCreateReq             = sys.LoginLogCreateReq
	LoginLogIndexReq              = sys.LoginLogIndexReq
	LoginLogIndexResp             = sys.LoginLogIndexResp
	LoginLogInfo                  = sys.LoginLogInfo
	Map                           = sys.Map
	MenuInfo                      = sys.MenuInfo
	MenuInfoIndexReq              = sys.MenuInfoIndexReq
	MenuInfoIndexResp             = sys.MenuInfoIndexResp
	OperLogCreateReq              = sys.OperLogCreateReq
	OperLogIndexReq               = sys.OperLogIndexReq
	OperLogIndexResp              = sys.OperLogIndexResp
	OperLogInfo                   = sys.OperLogInfo
	PageInfo                      = sys.PageInfo
	PageInfo_OrderBy              = sys.PageInfo_OrderBy
	Point                         = sys.Point
	ProjectInfo                   = sys.ProjectInfo
	ProjectInfoDeleteReq          = sys.ProjectInfoDeleteReq
	ProjectInfoIndexReq           = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp          = sys.ProjectInfoIndexResp
	ProjectInfoReadReq            = sys.ProjectInfoReadReq
	ReqWithID                     = sys.ReqWithID
	ReqWithIDCode                 = sys.ReqWithIDCode
	Response                      = sys.Response
	RoleApiAuthReq                = sys.RoleApiAuthReq
	RoleApiIndexReq               = sys.RoleApiIndexReq
	RoleApiIndexResp              = sys.RoleApiIndexResp
	RoleApiMultiUpdateReq         = sys.RoleApiMultiUpdateReq
	RoleAppIndexReq               = sys.RoleAppIndexReq
	RoleAppIndexResp              = sys.RoleAppIndexResp
	RoleAppMultiUpdateReq         = sys.RoleAppMultiUpdateReq
	RoleAppUpdateReq              = sys.RoleAppUpdateReq
	RoleInfo                      = sys.RoleInfo
	RoleInfoIndexReq              = sys.RoleInfoIndexReq
	RoleInfoIndexResp             = sys.RoleInfoIndexResp
	RoleMenuIndexReq              = sys.RoleMenuIndexReq
	RoleMenuIndexResp             = sys.RoleMenuIndexResp
	RoleMenuMultiUpdateReq        = sys.RoleMenuMultiUpdateReq
	TenantAppIndexReq             = sys.TenantAppIndexReq
	TenantAppIndexResp            = sys.TenantAppIndexResp
	TenantAppMultiUpdateReq       = sys.TenantAppMultiUpdateReq
	TenantInfo                    = sys.TenantInfo
	TenantInfoIndexReq            = sys.TenantInfoIndexReq
	TenantInfoIndexResp           = sys.TenantInfoIndexResp
	UserAuthArea                  = sys.UserAuthArea
	UserAuthAreaIndexReq          = sys.UserAuthAreaIndexReq
	UserAuthAreaIndexResp         = sys.UserAuthAreaIndexResp
	UserAuthAreaMultiUpdateReq    = sys.UserAuthAreaMultiUpdateReq
	UserAuthProject               = sys.UserAuthProject
	UserAuthProjectIndexReq       = sys.UserAuthProjectIndexReq
	UserAuthProjectIndexResp      = sys.UserAuthProjectIndexResp
	UserAuthProjectMultiUpdateReq = sys.UserAuthProjectMultiUpdateReq
	UserCheckTokenReq             = sys.UserCheckTokenReq
	UserCheckTokenResp            = sys.UserCheckTokenResp
	UserCreateResp                = sys.UserCreateResp
	UserInfo                      = sys.UserInfo
	UserInfoDeleteReq             = sys.UserInfoDeleteReq
	UserInfoIndexReq              = sys.UserInfoIndexReq
	UserInfoIndexResp             = sys.UserInfoIndexResp
	UserInfoReadReq               = sys.UserInfoReadReq
	UserLoginReq                  = sys.UserLoginReq
	UserLoginResp                 = sys.UserLoginResp
	UserRegister1Req              = sys.UserRegister1Req
	UserRegister1Resp             = sys.UserRegister1Resp
	UserRegister2Req              = sys.UserRegister2Req
	UserRoleIndexReq              = sys.UserRoleIndexReq
	UserRoleIndexResp             = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq        = sys.UserRoleMultiUpdateReq

	Common interface {
		Config(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ConfigResp, error)
	}

	defaultCommon struct {
		cli zrpc.Client
	}

	directCommon struct {
		svcCtx *svc.ServiceContext
		svr    sys.CommonServer
	}
)

func NewCommon(cli zrpc.Client) Common {
	return &defaultCommon{
		cli: cli,
	}
}

func NewDirectCommon(svcCtx *svc.ServiceContext, svr sys.CommonServer) Common {
	return &directCommon{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultCommon) Config(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ConfigResp, error) {
	client := sys.NewCommonClient(m.cli.Conn())
	return client.Config(ctx, in, opts...)
}

func (d *directCommon) Config(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ConfigResp, error) {
	return d.svr.Config(ctx, in)
}
