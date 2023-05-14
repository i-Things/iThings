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
	UserLoginSafeCtlResp  = sys.UserLoginSafeCtlResp
	UserReadReq           = sys.UserReadReq
	UserReadResp          = sys.UserReadResp
	UserUpdateReq         = sys.UserUpdateReq

	Auth interface {
		AuthApiCheck(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*Response, error)
		AuthApiMultiUpdate(ctx context.Context, in *AuthApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
		AuthApiIndex(ctx context.Context, in *AuthApiIndexReq, opts ...grpc.CallOption) (*AuthApiIndexResp, error)
	}

	defaultAuth struct {
		cli zrpc.Client
	}

	directAuth struct {
		svcCtx *svc.ServiceContext
		svr    sys.AuthServer
	}
)

func NewAuth(cli zrpc.Client) Auth {
	return &defaultAuth{
		cli: cli,
	}
}

func NewDirectAuth(svcCtx *svc.ServiceContext, svr sys.AuthServer) Auth {
	return &directAuth{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultAuth) AuthApiCheck(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewAuthClient(m.cli.Conn())
	return client.AuthApiCheck(ctx, in, opts...)
}

func (d *directAuth) AuthApiCheck(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.AuthApiCheck(ctx, in)
}

func (m *defaultAuth) AuthApiMultiUpdate(ctx context.Context, in *AuthApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewAuthClient(m.cli.Conn())
	return client.AuthApiMultiUpdate(ctx, in, opts...)
}

func (d *directAuth) AuthApiMultiUpdate(ctx context.Context, in *AuthApiMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.AuthApiMultiUpdate(ctx, in)
}

func (m *defaultAuth) AuthApiIndex(ctx context.Context, in *AuthApiIndexReq, opts ...grpc.CallOption) (*AuthApiIndexResp, error) {
	client := sys.NewAuthClient(m.cli.Conn())
	return client.AuthApiIndex(ctx, in, opts...)
}

func (d *directAuth) AuthApiIndex(ctx context.Context, in *AuthApiIndexReq, opts ...grpc.CallOption) (*AuthApiIndexResp, error) {
	return d.svr.AuthApiIndex(ctx, in)
}
