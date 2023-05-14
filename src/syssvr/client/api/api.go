// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiCreateReq          = sys.ApiCreateReq
	ApiData               = sys.ApiData
	ApiDeleteReq          = sys.ApiDeleteReq
	ApiIndexReq           = sys.ApiIndexReq
	ApiIndexResp          = sys.ApiIndexResp
	ApiUpdateReq          = sys.ApiUpdateReq
	AuthApiIndexReq       = sys.AuthApiIndexReq
	AuthApiIndexResp      = sys.AuthApiIndexResp
	AuthApiInfo           = sys.AuthApiInfo
	AuthApiMultiUpdateReq = sys.AuthApiMultiUpdateReq
	CheckAuthReq          = sys.CheckAuthReq
	ConfigResp            = sys.ConfigResp
	DateRange             = sys.DateRange
	JwtToken              = sys.JwtToken
	LoginLogCreateReq     = sys.LoginLogCreateReq
	LoginLogIndexData     = sys.LoginLogIndexData
	LoginLogIndexReq      = sys.LoginLogIndexReq
	LoginLogIndexResp     = sys.LoginLogIndexResp
	Map                   = sys.Map
	MenuCreateReq         = sys.MenuCreateReq
	MenuData              = sys.MenuData
	MenuDeleteReq         = sys.MenuDeleteReq
	MenuIndexReq          = sys.MenuIndexReq
	MenuIndexResp         = sys.MenuIndexResp
	MenuUpdateReq         = sys.MenuUpdateReq
	OperLogCreateReq      = sys.OperLogCreateReq
	OperLogIndexData      = sys.OperLogIndexData
	OperLogIndexReq       = sys.OperLogIndexReq
	OperLogIndexResp      = sys.OperLogIndexResp
	PageInfo              = sys.PageInfo
	Response              = sys.Response
	RoleCreateReq         = sys.RoleCreateReq
	RoleDeleteReq         = sys.RoleDeleteReq
	RoleIndexData         = sys.RoleIndexData
	RoleIndexReq          = sys.RoleIndexReq
	RoleIndexResp         = sys.RoleIndexResp
	RoleMenuUpdateReq     = sys.RoleMenuUpdateReq
	RoleUpdateReq         = sys.RoleUpdateReq
	UserCheckTokenReq     = sys.UserCheckTokenReq
	UserCheckTokenResp    = sys.UserCheckTokenResp
	UserCreateReq         = sys.UserCreateReq
	UserCreateResp        = sys.UserCreateResp
	UserDeleteReq         = sys.UserDeleteReq
	UserIndexReq          = sys.UserIndexReq
	UserIndexResp         = sys.UserIndexResp
	UserInfo              = sys.UserInfo
	UserLoginReq          = sys.UserLoginReq
	UserLoginResp         = sys.UserLoginResp
	UserLoginSafeCtlReq   = sys.UserLoginSafeCtlReq
	UserReadReq           = sys.UserReadReq
	UserReadResp          = sys.UserReadResp
	UserUpdateReq         = sys.UserUpdateReq

	Api interface {
		ApiCreate(ctx context.Context, in *ApiCreateReq, opts ...grpc.CallOption) (*Response, error)
		ApiIndex(ctx context.Context, in *ApiIndexReq, opts ...grpc.CallOption) (*ApiIndexResp, error)
		ApiUpdate(ctx context.Context, in *ApiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		ApiDelete(ctx context.Context, in *ApiDeleteReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultApi struct {
		cli zrpc.Client
	}

	directApi struct {
		svcCtx *svc.ServiceContext
		svr    sys.ApiServer
	}
)

func NewApi(cli zrpc.Client) Api {
	return &defaultApi{
		cli: cli,
	}
}

func NewDirectApi(svcCtx *svc.ServiceContext, svr sys.ApiServer) Api {
	return &directApi{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultApi) ApiCreate(ctx context.Context, in *ApiCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewApiClient(m.cli.Conn())
	return client.ApiCreate(ctx, in, opts...)
}

func (d *directApi) ApiCreate(ctx context.Context, in *ApiCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ApiCreate(ctx, in)
}

func (m *defaultApi) ApiIndex(ctx context.Context, in *ApiIndexReq, opts ...grpc.CallOption) (*ApiIndexResp, error) {
	client := sys.NewApiClient(m.cli.Conn())
	return client.ApiIndex(ctx, in, opts...)
}

func (d *directApi) ApiIndex(ctx context.Context, in *ApiIndexReq, opts ...grpc.CallOption) (*ApiIndexResp, error) {
	return d.svr.ApiIndex(ctx, in)
}

func (m *defaultApi) ApiUpdate(ctx context.Context, in *ApiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewApiClient(m.cli.Conn())
	return client.ApiUpdate(ctx, in, opts...)
}

func (d *directApi) ApiUpdate(ctx context.Context, in *ApiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ApiUpdate(ctx, in)
}

func (m *defaultApi) ApiDelete(ctx context.Context, in *ApiDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewApiClient(m.cli.Conn())
	return client.ApiDelete(ctx, in, opts...)
}

func (d *directApi) ApiDelete(ctx context.Context, in *ApiDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ApiDelete(ctx, in)
}
