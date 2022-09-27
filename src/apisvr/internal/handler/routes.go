// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	systemmenu "github.com/i-Things/things/src/apisvr/internal/handler/system/menu"
	systemrole "github.com/i-Things/things/src/apisvr/internal/handler/system/role"
	systemuser "github.com/i-Things/things/src/apisvr/internal/handler/system/user"
	thingsdeviceauth "github.com/i-Things/things/src/apisvr/internal/handler/things/device/auth"
	thingsdeviceinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/device/info"
	thingsdeviceinteract "github.com/i-Things/things/src/apisvr/internal/handler/things/device/interact"
	thingsdevicemsg "github.com/i-Things/things/src/apisvr/internal/handler/things/device/msg"
	thingsgroupdevice "github.com/i-Things/things/src/apisvr/internal/handler/things/group/device"
	thingsgroupinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/group/info"
	thingsproductinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/product/info"
	thingsproductschema "github.com/i-Things/things/src/apisvr/internal/handler/things/product/schema"
	"github.com/i-Things/things/src/apisvr/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/captcha",
				Handler: systemuser.CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: systemuser.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemuser.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemuser.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemuser.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: systemuser.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemuser.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/resource-read",
					Handler: systemuser.ResourceReadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemmenu.MenuCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemmenu.MenuIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemmenu.MenuUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemmenu.MenuDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/menu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemrole.RoleCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemrole.RoleIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemrole.RoleUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemrole.RoleDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role-menu/update",
					Handler: systemrole.RoleMenuUpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/role"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: thingsdeviceauth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/access",
				Handler: thingsdeviceauth.AccessHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/root-check",
				Handler: thingsdeviceauth.RootCheckHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/hub-log/index",
				Handler: thingsdevicemsg.HubLogIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sdk-log/index",
				Handler: thingsdevicemsg.SdkLogIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/property-log/index",
				Handler: thingsdevicemsg.PropertyLogIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/property-latest/index",
				Handler: thingsdevicemsg.PropertyLatestIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/event-log/index",
				Handler: thingsdevicemsg.EventLogIndexHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/msg"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: thingsdeviceinfo.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: thingsdeviceinfo.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: thingsdeviceinfo.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/index",
				Handler: thingsdeviceinfo.IndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/read",
				Handler: thingsdeviceinfo.ReadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/info"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/send-action",
				Handler: thingsdeviceinteract.SendActionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/send-property",
				Handler: thingsdeviceinteract.SendPropertyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/send-msg",
				Handler: thingsdeviceinteract.SendMsgHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/interact"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: thingsproductinfo.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: thingsproductinfo.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: thingsproductinfo.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/index",
				Handler: thingsproductinfo.IndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/read",
				Handler: thingsproductinfo.ReadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/product/info"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/tsl-import",
				Handler: thingsproductschema.TslImportHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/tsl-read",
				Handler: thingsproductschema.TslReadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: thingsproductschema.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: thingsproductschema.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: thingsproductschema.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/index",
				Handler: thingsproductschema.IndexHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/product/schema"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsgroupinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsgroupinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsgroupinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsgroupinfo.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsgroupinfo.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/group/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckToken},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsgroupdevice.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsgroupdevice.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsgroupdevice.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/group/device"),
	)
}
