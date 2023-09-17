package deviceDelete

import (
	"context"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

func DeviceGroupHandle(svcCtx *svc.ServiceContext) any {
	return func(ctx context.Context, value *devices.Core) {
		err := relationDB.NewGroupDeviceRepo(ctx).DeleteByFilter(ctx, relationDB.GroupDeviceFilter{
			ProductID:  value.ProductID,
			DeviceName: value.DeviceName,
		})
		logx.WithContext(ctx).Infof("DeviceGroupHandle value:%v err:%v", utils.Fmt(value), err)
	}
}
