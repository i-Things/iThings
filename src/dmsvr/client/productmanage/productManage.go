// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package productmanage

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessAuthReq               = dm.AccessAuthReq
	DeleteOtaFirmwareReq        = dm.DeleteOtaFirmwareReq
	DeviceCore                  = dm.DeviceCore
	DeviceGatewayBindDevice     = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq       = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp      = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign           = dm.DeviceGatewaySign
	DeviceInfo                  = dm.DeviceInfo
	DeviceInfoCountReq          = dm.DeviceInfoCountReq
	DeviceInfoCountResp         = dm.DeviceInfoCountResp
	DeviceInfoDeleteReq         = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq          = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp         = dm.DeviceInfoIndexResp
	DeviceInfoReadReq           = dm.DeviceInfoReadReq
	DeviceRegisterReq           = dm.DeviceRegisterReq
	DeviceRegisterResp          = dm.DeviceRegisterResp
	DeviceTypeCountReq          = dm.DeviceTypeCountReq
	DeviceTypeCountResp         = dm.DeviceTypeCountResp
	EventIndex                  = dm.EventIndex
	EventIndexResp              = dm.EventIndexResp
	EventLogIndexReq            = dm.EventLogIndexReq
	Firmware                    = dm.Firmware
	FirmwareFile                = dm.FirmwareFile
	FirmwareInfo                = dm.FirmwareInfo
	FirmwareInfoDeleteReq       = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp      = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq        = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp       = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq         = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp        = dm.FirmwareInfoReadResp
	FirmwareResp                = dm.FirmwareResp
	GetPropertyReplyReq         = dm.GetPropertyReplyReq
	GetPropertyReplyResp        = dm.GetPropertyReplyResp
	GroupDeviceIndexReq         = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp        = dm.GroupDeviceIndexResp
	GroupDeviceMultiCreateReq   = dm.GroupDeviceMultiCreateReq
	GroupDeviceMultiDeleteReq   = dm.GroupDeviceMultiDeleteReq
	GroupInfo                   = dm.GroupInfo
	GroupInfoCreateReq          = dm.GroupInfoCreateReq
	GroupInfoDeleteReq          = dm.GroupInfoDeleteReq
	GroupInfoIndexReq           = dm.GroupInfoIndexReq
	GroupInfoIndexResp          = dm.GroupInfoIndexResp
	GroupInfoReadReq            = dm.GroupInfoReadReq
	GroupInfoUpdateReq          = dm.GroupInfoUpdateReq
	HubLogIndex                 = dm.HubLogIndex
	HubLogIndexReq              = dm.HubLogIndexReq
	HubLogIndexResp             = dm.HubLogIndexResp
	ListOtaFirmwareReq          = dm.ListOtaFirmwareReq
	ListOtaFirmwareResp         = dm.ListOtaFirmwareResp
	LoginAuthReq                = dm.LoginAuthReq
	ModifyOtaFirmwareReq        = dm.ModifyOtaFirmwareReq
	MultiSendPropertyReq        = dm.MultiSendPropertyReq
	MultiSendPropertyResp       = dm.MultiSendPropertyResp
	OtaCommonResp               = dm.OtaCommonResp
	OtaFirmwareDeviceInfoReq    = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp   = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile             = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq     = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp    = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo         = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq          = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp         = dm.OtaFirmwareFileResp
	OtaFirmwareInfo             = dm.OtaFirmwareInfo
	OtaFirmwareReq              = dm.OtaFirmwareReq
	OtaFirmwareResp             = dm.OtaFirmwareResp
	OtaPageInfo                 = dm.OtaPageInfo
	OtaPromptIndexReq           = dm.OtaPromptIndexReq
	OtaPromptIndexResp          = dm.OtaPromptIndexResp
	OtaTaskBatchReq             = dm.OtaTaskBatchReq
	OtaTaskBatchResp            = dm.OtaTaskBatchResp
	OtaTaskCancleReq            = dm.OtaTaskCancleReq
	OtaTaskCreatResp            = dm.OtaTaskCreatResp
	OtaTaskCreateReq            = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq      = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq       = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp      = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo           = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq     = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq        = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq             = dm.OtaTaskIndexReq
	OtaTaskIndexResp            = dm.OtaTaskIndexResp
	OtaTaskInfo                 = dm.OtaTaskInfo
	OtaTaskReadReq              = dm.OtaTaskReadReq
	OtaTaskReadResp             = dm.OtaTaskReadResp
	PageInfo                    = dm.PageInfo
	PageInfo_OrderBy            = dm.PageInfo_OrderBy
	Point                       = dm.Point
	ProductCustom               = dm.ProductCustom
	ProductCustomReadReq        = dm.ProductCustomReadReq
	ProductInfo                 = dm.ProductInfo
	ProductInfoDeleteReq        = dm.ProductInfoDeleteReq
	ProductInfoIndexReq         = dm.ProductInfoIndexReq
	ProductInfoIndexResp        = dm.ProductInfoIndexResp
	ProductInfoReadReq          = dm.ProductInfoReadReq
	ProductRemoteConfig         = dm.ProductRemoteConfig
	ProductSchemaCreateReq      = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq      = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq       = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp      = dm.ProductSchemaIndexResp
	ProductSchemaInfo           = dm.ProductSchemaInfo
	ProductSchemaTslImportReq   = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq     = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp    = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq      = dm.ProductSchemaUpdateReq
	PropertyIndex               = dm.PropertyIndex
	PropertyIndexResp           = dm.PropertyIndexResp
	PropertyLatestIndexReq      = dm.PropertyLatestIndexReq
	PropertyLogIndexReq         = dm.PropertyLogIndexReq
	QueryOtaFirmwareReq         = dm.QueryOtaFirmwareReq
	QueryOtaFirmwareResp        = dm.QueryOtaFirmwareResp
	RemoteConfigCreateReq       = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq        = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp       = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq     = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp    = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq      = dm.RemoteConfigPushAllReq
	RespActionReq               = dm.RespActionReq
	RespReadReq                 = dm.RespReadReq
	Response                    = dm.Response
	RootCheckReq                = dm.RootCheckReq
	SdkLogIndex                 = dm.SdkLogIndex
	SdkLogIndexReq              = dm.SdkLogIndexReq
	SdkLogIndexResp             = dm.SdkLogIndexResp
	SendActionReq               = dm.SendActionReq
	SendActionResp              = dm.SendActionResp
	SendMsgReq                  = dm.SendMsgReq
	SendMsgResp                 = dm.SendMsgResp
	SendOption                  = dm.SendOption
	SendPropertyMsg             = dm.SendPropertyMsg
	SendPropertyReq             = dm.SendPropertyReq
	SendPropertyResp            = dm.SendPropertyResp
	ShadowIndex                 = dm.ShadowIndex
	ShadowIndexResp             = dm.ShadowIndexResp
	Tag                         = dm.Tag
	VerifyOtaFirmwareReq        = dm.VerifyOtaFirmwareReq

	ProductManage interface {
		// 新增产品
		ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
		// 更新产品
		ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除产品
		ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取产品信息列表
		ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error)
		// 获取产品信息详情
		ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error)
		// 更新产品物模型
		ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error)
		// 新增产品
		ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Response, error)
		// 删除产品
		ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取产品信息列表
		ProductSchemaIndex(ctx context.Context, in *ProductSchemaIndexReq, opts ...grpc.CallOption) (*ProductSchemaIndexResp, error)
		// 删除产品
		ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Response, error)
		// 获取产品信息列表
		ProductSchemaTslRead(ctx context.Context, in *ProductSchemaTslReadReq, opts ...grpc.CallOption) (*ProductSchemaTslReadResp, error)
		// 脚本管理
		ProductCustomRead(ctx context.Context, in *ProductCustomReadReq, opts ...grpc.CallOption) (*ProductCustom, error)
		ProductCustomUpdate(ctx context.Context, in *ProductCustom, opts ...grpc.CallOption) (*Response, error)
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

// 新增产品
func (m *defaultProductManage) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoCreate(ctx, in, opts...)
}

// 新增产品
func (d *directProductManage) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoCreate(ctx, in)
}

// 更新产品
func (m *defaultProductManage) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoUpdate(ctx, in, opts...)
}

// 更新产品
func (d *directProductManage) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductInfoUpdate(ctx, in)
}

// 删除产品
func (m *defaultProductManage) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductInfoDelete(ctx, in, opts...)
}

// 删除产品
func (d *directProductManage) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
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
func (m *defaultProductManage) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaUpdate(ctx, in, opts...)
}

// 更新产品物模型
func (d *directProductManage) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductSchemaUpdate(ctx, in)
}

// 新增产品
func (m *defaultProductManage) ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaCreate(ctx, in, opts...)
}

// 新增产品
func (d *directProductManage) ProductSchemaCreate(ctx context.Context, in *ProductSchemaCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductSchemaCreate(ctx, in)
}

// 删除产品
func (m *defaultProductManage) ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaDelete(ctx, in, opts...)
}

// 删除产品
func (d *directProductManage) ProductSchemaDelete(ctx context.Context, in *ProductSchemaDeleteReq, opts ...grpc.CallOption) (*Response, error) {
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
func (m *defaultProductManage) ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductSchemaTslImport(ctx, in, opts...)
}

// 删除产品
func (d *directProductManage) ProductSchemaTslImport(ctx context.Context, in *ProductSchemaTslImportReq, opts ...grpc.CallOption) (*Response, error) {
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

func (m *defaultProductManage) ProductCustomUpdate(ctx context.Context, in *ProductCustom, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewProductManageClient(m.cli.Conn())
	return client.ProductCustomUpdate(ctx, in, opts...)
}

func (d *directProductManage) ProductCustomUpdate(ctx context.Context, in *ProductCustom, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.ProductCustomUpdate(ctx, in)
}
