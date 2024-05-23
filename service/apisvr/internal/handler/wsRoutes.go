// Code generated by goctl. DO NOT EDIT.
package handler

import (
	ws "gitee.com/i-Things/share/websocket"
	"net/http"

	thingsdeviceauth "github.com/i-Things/things/service/apisvr/internal/handler/things/device/auth"
	thingsdeviceauth5 "github.com/i-Things/things/service/apisvr/internal/handler/things/device/auth5"
	thingsdevicegateway "github.com/i-Things/things/service/apisvr/internal/handler/things/device/gateway"
	thingsdeviceinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/device/info"
	thingsdeviceinteract "github.com/i-Things/things/service/apisvr/internal/handler/things/device/interact"
	thingsdevicemoduleversion "github.com/i-Things/things/service/apisvr/internal/handler/things/device/module/version"
	thingsdevicemsg "github.com/i-Things/things/service/apisvr/internal/handler/things/device/msg"
	thingsdeviceprofile "github.com/i-Things/things/service/apisvr/internal/handler/things/device/profile"
	thingsgroupdevice "github.com/i-Things/things/service/apisvr/internal/handler/things/group/device"
	thingsgroupinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/group/info"
	thingsotafirmwaredevice "github.com/i-Things/things/service/apisvr/internal/handler/things/ota/firmware/device"
	thingsotafirmwareinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/ota/firmware/info"
	thingsotafirmwarejob "github.com/i-Things/things/service/apisvr/internal/handler/things/ota/firmware/job"
	thingsotamoduleinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/ota/module/info"
	thingsproductcategory "github.com/i-Things/things/service/apisvr/internal/handler/things/product/category"
	thingsproductcustom "github.com/i-Things/things/service/apisvr/internal/handler/things/product/custom"
	thingsproductinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/product/info"
	thingsproductremoteConfig "github.com/i-Things/things/service/apisvr/internal/handler/things/product/remoteConfig"
	thingsproductschema "github.com/i-Things/things/service/apisvr/internal/handler/things/product/schema"
	thingsprotocolinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/protocol/info"
	thingsruledeviceTimerinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/rule/deviceTimer/info"
	thingsrulesceneinfo "github.com/i-Things/things/service/apisvr/internal/handler/things/rule/scene/info"
	thingsschemacommon "github.com/i-Things/things/service/apisvr/internal/handler/things/schema/common"
	thingsslotarea "github.com/i-Things/things/service/apisvr/internal/handler/things/slot/area"
	thingsslotuser "github.com/i-Things/things/service/apisvr/internal/handler/things/slot/user"
	thingsuserdevicecollect "github.com/i-Things/things/service/apisvr/internal/handler/things/user/device/collect"
	thingsuserdeviceshare "github.com/i-Things/things/service/apisvr/internal/handler/things/user/device/share"
	"github.com/i-Things/things/service/apisvr/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterWsHandlers(server *ws.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.InitCtxsWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/access",
					Handler: thingsdeviceauth.AccessHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: thingsdeviceauth.LoginHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/register",
					Handler: thingsdeviceauth.RegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/root-check",
					Handler: thingsdeviceauth.RootCheckHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/auth"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.InitCtxsWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/access",
					Handler: thingsdeviceauth5.AccessHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/login",
					Handler: thingsdeviceauth5.LoginHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/auth5"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsdevicegateway.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-create",
					Handler: thingsdevicegateway.MultiCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-delete",
					Handler: thingsdevicegateway.MultiDeleteHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/gateway"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/bind",
					Handler: thingsdeviceinfo.BindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/canBind",
					Handler: thingsdeviceinfo.CanBindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/count",
					Handler: thingsdeviceinfo.CountHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsdeviceinfo.CreateHandler(serverCtx),
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
					Path:    "/multi-import",
					Handler: thingsdeviceinfo.MultiImportHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-update",
					Handler: thingsdeviceinfo.MultiUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsdeviceinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/transfer",
					Handler: thingsdeviceinfo.TransferHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/unbind",
					Handler: thingsdeviceinfo.UnbindHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsdeviceinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action-read",
					Handler: thingsdeviceinteract.ActionReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/action-send",
					Handler: thingsdeviceinteract.ActionSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/gateway-get-found-send",
					Handler: thingsdeviceinteract.GatewayGetFoundSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/gateway-notify-bind-send",
					Handler: thingsdeviceinteract.GatewayNotifyBindSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-send-property",
					Handler: thingsdeviceinteract.MultiSendPropertyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-control-multi-send",
					Handler: thingsdeviceinteract.PropertyControlMultiSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-control-read",
					Handler: thingsdeviceinteract.PropertyControlReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-control-send",
					Handler: thingsdeviceinteract.PropertyControlSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-get-report-send",
					Handler: thingsdeviceinteract.PropertyGetReportSendHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-read",
					Handler: thingsdeviceinteract.PropertyReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/send-action",
					Handler: thingsdeviceinteract.SendActionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/send-msg",
					Handler: thingsdeviceinteract.SendMsgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/send-property",
					Handler: thingsdeviceinteract.SendPropertyHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/interact"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsdevicemoduleversion.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsdevicemoduleversion.ReadHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/version"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/event-log/index",
					Handler: thingsdevicemsg.EventLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/gateway-can-bind-index",
					Handler: thingsdevicemsg.GatewayCanBindIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/hub-log/index",
					Handler: thingsdevicemsg.HubLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-log-latest/index",
					Handler: thingsdevicemsg.PropertyLatestIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-log/index",
					Handler: thingsdevicemsg.PropertyLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sdk-log/index",
					Handler: thingsdevicemsg.SdkLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/send-log/index",
					Handler: thingsdevicemsg.SendLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/shadow/index",
					Handler: thingsdevicemsg.ShadowIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/status-log/index",
					Handler: thingsdevicemsg.StatusLogIndexHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/msg"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsdeviceprofile.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsdeviceprofile.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsdeviceprofile.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsdeviceprofile.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/device/profile"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsgroupdevice.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-create",
					Handler: thingsgroupdevice.MultiCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-delete",
					Handler: thingsgroupdevice.MultiDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-update",
					Handler: thingsgroupdevice.MultiUpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/group/device"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsgroupinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsgroupinfo.DeleteHandler(serverCtx),
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
			}...,
		),
		ws.WithPrefix("/api/v1/things/group/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/cancel",
					Handler: thingsotafirmwaredevice.CancelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsotafirmwaredevice.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/retry",
					Handler: thingsotafirmwaredevice.RetryHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/ota/firmware/device"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsotafirmwareinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsotafirmwareinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsotafirmwareinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsotafirmwareinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsotafirmwareinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/ota/firmware/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsotafirmwarejob.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsotafirmwarejob.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsotafirmwarejob.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsotafirmwarejob.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/ota/firmware/job"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsotamoduleinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsotamoduleinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsotamoduleinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsotamoduleinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsotamoduleinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/ota/module/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsproductcategory.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsproductcategory.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsproductcategory.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsproductcategory.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema/index",
					Handler: thingsproductcategory.SchemaIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema/multi-create",
					Handler: thingsproductcategory.SchemaMultiCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema/multi-delete",
					Handler: thingsproductcategory.SchemaMultiDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/schema/multi-update",
					Handler: thingsproductcategory.SchemaMultiUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsproductcategory.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/product/category"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsproductcustom.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsproductcustom.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/product/custom"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsproductinfo.CreateHandler(serverCtx),
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
					Path:    "/init",
					Handler: thingsproductinfo.InitHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsproductinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsproductinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/product/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/remote-config/create",
					Handler: thingsproductremoteConfig.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/remote-config/index",
					Handler: thingsproductremoteConfig.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/remote-config/lastest-read",
					Handler: thingsproductremoteConfig.LastestReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/remote-config/push-all",
					Handler: thingsproductremoteConfig.PushAllHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
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
				{
					Method:  http.MethodPost,
					Path:    "/multi-create",
					Handler: thingsproductschema.MultiCreateHandler(serverCtx),
				},
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
			}...,
		),
		ws.WithPrefix("/api/v1/things/product/schema"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsprotocolinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsprotocolinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsprotocolinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsprotocolinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsprotocolinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/protocol/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/device-timer/info/create",
					Handler: thingsruledeviceTimerinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device-timer/info/delete",
					Handler: thingsruledeviceTimerinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device-timer/info/index",
					Handler: thingsruledeviceTimerinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device-timer/info/read",
					Handler: thingsruledeviceTimerinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/device-timer/info/update",
					Handler: thingsruledeviceTimerinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/rule"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsrulesceneinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsrulesceneinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsrulesceneinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/manually-trigger",
					Handler: thingsrulesceneinfo.ManuallyTriggerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsrulesceneinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsrulesceneinfo.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/rule/scene/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.CheckApiWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsschemacommon.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsschemacommon.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsschemacommon.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/init",
					Handler: thingsschemacommon.InitHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsschemacommon.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/schema/common"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsslotarea.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsslotarea.DeleteHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/slot/area"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/subscribe",
					Handler: thingsslotuser.SubscribeHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/slot/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsuserdevicecollect.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-create",
					Handler: thingsuserdevicecollect.MultiCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-delete",
					Handler: thingsuserdevicecollect.MultiDeleteHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/user/device/collect"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckTokenWare, serverCtx.InitCtxsWare, serverCtx.DataAuthWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsuserdeviceshare.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsuserdeviceshare.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsuserdeviceshare.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-delete",
					Handler: thingsuserdeviceshare.MultiDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsuserdeviceshare.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsuserdeviceshare.UpdateHandler(serverCtx),
				},
			}...,
		),
		ws.WithPrefix("/api/v1/things/user/device/share"),
	)
}
