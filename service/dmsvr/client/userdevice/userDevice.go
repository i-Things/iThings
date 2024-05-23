// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package userdevice

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
	ManufacturerInfo                  = dm.ManufacturerInfo
	OtaFirmwareDeviceCancelReq        = dm.OtaFirmwareDeviceCancelReq
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
	ProductCategoryReadReq            = dm.ProductCategoryReadReq
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
	WithIDCode                        = dm.WithIDCode

	UserDevice interface {
		// 用户收藏的设备
		UserDeviceCollectMultiCreate(ctx context.Context, in *UserDeviceCollectSave, opts ...grpc.CallOption) (*Empty, error)
		UserDeviceCollectMultiDelete(ctx context.Context, in *UserDeviceCollectSave, opts ...grpc.CallOption) (*Empty, error)
		UserDeviceCollectIndex(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserDeviceCollectSave, error)
		// 分享设备
		UserDeviceShareCreate(ctx context.Context, in *UserDeviceShareInfo, opts ...grpc.CallOption) (*WithID, error)
		// 更新权限
		UserDeviceShareUpdate(ctx context.Context, in *UserDeviceShareInfo, opts ...grpc.CallOption) (*Empty, error)
		// 取消分享设备
		UserDeviceShareDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		// 取消分享设备
		UserDeviceShareMultiDelete(ctx context.Context, in *UserDeviceShareMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取设备分享列表(只有设备的所有者才能获取)
		UserDeviceShareIndex(ctx context.Context, in *UserDeviceShareIndexReq, opts ...grpc.CallOption) (*UserDeviceShareIndexResp, error)
		// 获取设备分享的详情
		UserDeviceShareRead(ctx context.Context, in *UserDeviceShareReadReq, opts ...grpc.CallOption) (*UserDeviceShareInfo, error)
		UserDeviceTransfer(ctx context.Context, in *DeviceTransferReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultUserDevice struct {
		cli zrpc.Client
	}

	directUserDevice struct {
		svcCtx *svc.ServiceContext
		svr    dm.UserDeviceServer
	}
)

func NewUserDevice(cli zrpc.Client) UserDevice {
	return &defaultUserDevice{
		cli: cli,
	}
}

func NewDirectUserDevice(svcCtx *svc.ServiceContext, svr dm.UserDeviceServer) UserDevice {
	return &directUserDevice{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 用户收藏的设备
func (m *defaultUserDevice) UserDeviceCollectMultiCreate(ctx context.Context, in *UserDeviceCollectSave, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceCollectMultiCreate(ctx, in, opts...)
}

// 用户收藏的设备
func (d *directUserDevice) UserDeviceCollectMultiCreate(ctx context.Context, in *UserDeviceCollectSave, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.UserDeviceCollectMultiCreate(ctx, in)
}

func (m *defaultUserDevice) UserDeviceCollectMultiDelete(ctx context.Context, in *UserDeviceCollectSave, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceCollectMultiDelete(ctx, in, opts...)
}

func (d *directUserDevice) UserDeviceCollectMultiDelete(ctx context.Context, in *UserDeviceCollectSave, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.UserDeviceCollectMultiDelete(ctx, in)
}

func (m *defaultUserDevice) UserDeviceCollectIndex(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserDeviceCollectSave, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceCollectIndex(ctx, in, opts...)
}

func (d *directUserDevice) UserDeviceCollectIndex(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserDeviceCollectSave, error) {
	return d.svr.UserDeviceCollectIndex(ctx, in)
}

// 分享设备
func (m *defaultUserDevice) UserDeviceShareCreate(ctx context.Context, in *UserDeviceShareInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceShareCreate(ctx, in, opts...)
}

// 分享设备
func (d *directUserDevice) UserDeviceShareCreate(ctx context.Context, in *UserDeviceShareInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.UserDeviceShareCreate(ctx, in)
}

// 更新权限
func (m *defaultUserDevice) UserDeviceShareUpdate(ctx context.Context, in *UserDeviceShareInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceShareUpdate(ctx, in, opts...)
}

// 更新权限
func (d *directUserDevice) UserDeviceShareUpdate(ctx context.Context, in *UserDeviceShareInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.UserDeviceShareUpdate(ctx, in)
}

// 取消分享设备
func (m *defaultUserDevice) UserDeviceShareDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceShareDelete(ctx, in, opts...)
}

// 取消分享设备
func (d *directUserDevice) UserDeviceShareDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.UserDeviceShareDelete(ctx, in)
}

// 取消分享设备
func (m *defaultUserDevice) UserDeviceShareMultiDelete(ctx context.Context, in *UserDeviceShareMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceShareMultiDelete(ctx, in, opts...)
}

// 取消分享设备
func (d *directUserDevice) UserDeviceShareMultiDelete(ctx context.Context, in *UserDeviceShareMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.UserDeviceShareMultiDelete(ctx, in)
}

// 获取设备分享列表(只有设备的所有者才能获取)
func (m *defaultUserDevice) UserDeviceShareIndex(ctx context.Context, in *UserDeviceShareIndexReq, opts ...grpc.CallOption) (*UserDeviceShareIndexResp, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceShareIndex(ctx, in, opts...)
}

// 获取设备分享列表(只有设备的所有者才能获取)
func (d *directUserDevice) UserDeviceShareIndex(ctx context.Context, in *UserDeviceShareIndexReq, opts ...grpc.CallOption) (*UserDeviceShareIndexResp, error) {
	return d.svr.UserDeviceShareIndex(ctx, in)
}

// 获取设备分享的详情
func (m *defaultUserDevice) UserDeviceShareRead(ctx context.Context, in *UserDeviceShareReadReq, opts ...grpc.CallOption) (*UserDeviceShareInfo, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceShareRead(ctx, in, opts...)
}

// 获取设备分享的详情
func (d *directUserDevice) UserDeviceShareRead(ctx context.Context, in *UserDeviceShareReadReq, opts ...grpc.CallOption) (*UserDeviceShareInfo, error) {
	return d.svr.UserDeviceShareRead(ctx, in)
}

func (m *defaultUserDevice) UserDeviceTransfer(ctx context.Context, in *DeviceTransferReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewUserDeviceClient(m.cli.Conn())
	return client.UserDeviceTransfer(ctx, in, opts...)
}

func (d *directUserDevice) UserDeviceTransfer(ctx context.Context, in *DeviceTransferReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.UserDeviceTransfer(ctx, in)
}
