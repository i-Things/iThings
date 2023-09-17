package devicemanagelogic

import (
	"context"
	"fmt"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceInfoCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	DiDB *relationDB.DeviceInfoRepo
}

func NewDeviceInfoCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceInfoCountLogic {
	return &DeviceInfoCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		DiDB:   relationDB.NewDeviceInfoRepo(ctx),
	}
}

// 设备计数
func (l *DeviceInfoCountLogic) DeviceInfoCount(in *dm.DeviceInfoCountReq) (*dm.DeviceInfoCountResp, error) {
	diCount, err := l.DiDB.CountGroupByField(
		l.ctx,
		relationDB.DeviceFilter{
			LastLoginTime: struct {
				Start int64
				End   int64
			}{Start: in.StartTime, End: in.EndTime},
		},
		"is_online")
	if err != nil {
		if errors.Cmp(err, errors.NotFind) {
			return nil, errors.NotFind
		}
		return nil, err
	}

	onlineCount := diCount[fmt.Sprintf("%d", def.DeviceStatusOnline)]
	offlineCount := diCount[fmt.Sprintf("%d", def.DeviceStatusOffline)]
	InactiveCount := diCount[fmt.Sprintf("%d", def.DeviceStatusInactive)]
	var allCount int64
	for _, v := range diCount {
		allCount += v
	}

	return &dm.DeviceInfoCountResp{
		Online:   onlineCount,
		Offline:  offlineCount,
		Inactive: InactiveCount,
		Unknown:  allCount - onlineCount - offlineCount - InactiveCount,
	}, nil
}
