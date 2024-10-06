// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package server

import (
	"context"

	"gitee.com/i-Things/things/service/dmsvr/internal/logic/devicemsg"
	"gitee.com/i-Things/things/service/dmsvr/internal/svc"
	"gitee.com/i-Things/things/service/dmsvr/pb/dm"
)

type DeviceMsgServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedDeviceMsgServer
}

func NewDeviceMsgServer(svcCtx *svc.ServiceContext) *DeviceMsgServer {
	return &DeviceMsgServer{
		svcCtx: svcCtx,
	}
}

// 获取设备sdk调试日志
func (s *DeviceMsgServer) SdkLogIndex(ctx context.Context, in *dm.SdkLogIndexReq) (*dm.SdkLogIndexResp, error) {
	l := devicemsglogic.NewSdkLogIndexLogic(ctx, s.svcCtx)
	return l.SdkLogIndex(in)
}

// 获取设备调试信息记录登入登出,操作
func (s *DeviceMsgServer) HubLogIndex(ctx context.Context, in *dm.HubLogIndexReq) (*dm.HubLogIndexResp, error) {
	l := devicemsglogic.NewHubLogIndexLogic(ctx, s.svcCtx)
	return l.HubLogIndex(in)
}

func (s *DeviceMsgServer) SendLogIndex(ctx context.Context, in *dm.SendLogIndexReq) (*dm.SendLogIndexResp, error) {
	l := devicemsglogic.NewSendLogIndexLogic(ctx, s.svcCtx)
	return l.SendLogIndex(in)
}

func (s *DeviceMsgServer) StatusLogIndex(ctx context.Context, in *dm.StatusLogIndexReq) (*dm.StatusLogIndexResp, error) {
	l := devicemsglogic.NewStatusLogIndexLogic(ctx, s.svcCtx)
	return l.StatusLogIndex(in)
}

// 获取设备数据信息
func (s *DeviceMsgServer) PropertyLogLatestIndex(ctx context.Context, in *dm.PropertyLogLatestIndexReq) (*dm.PropertyLogIndexResp, error) {
	l := devicemsglogic.NewPropertyLogLatestIndexLogic(ctx, s.svcCtx)
	return l.PropertyLogLatestIndex(in)
}

// 获取设备数据信息
func (s *DeviceMsgServer) PropertyLogIndex(ctx context.Context, in *dm.PropertyLogIndexReq) (*dm.PropertyLogIndexResp, error) {
	l := devicemsglogic.NewPropertyLogIndexLogic(ctx, s.svcCtx)
	return l.PropertyLogIndex(in)
}

// 获取设备数据信息
func (s *DeviceMsgServer) EventLogIndex(ctx context.Context, in *dm.EventLogIndexReq) (*dm.EventLogIndexResp, error) {
	l := devicemsglogic.NewEventLogIndexLogic(ctx, s.svcCtx)
	return l.EventLogIndex(in)
}

// 获取设备影子列表
func (s *DeviceMsgServer) ShadowIndex(ctx context.Context, in *dm.PropertyLogLatestIndexReq) (*dm.ShadowIndexResp, error) {
	l := devicemsglogic.NewShadowIndexLogic(ctx, s.svcCtx)
	return l.ShadowIndex(in)
}

// 获取网关可以绑定的子设备列表
func (s *DeviceMsgServer) GatewayCanBindIndex(ctx context.Context, in *dm.GatewayCanBindIndexReq) (*dm.GatewayCanBindIndexResp, error) {
	l := devicemsglogic.NewGatewayCanBindIndexLogic(ctx, s.svcCtx)
	return l.GatewayCanBindIndex(in)
}
