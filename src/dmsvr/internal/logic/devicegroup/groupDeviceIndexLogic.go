package devicegrouplogic

import (
	"context"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/logic"
	"github.com/i-Things/things/src/dmsvr/internal/logic/devicemanage"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupDeviceIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GdDB *relationDB.GroupDeviceRepo
	DiDB *relationDB.DeviceInfoRepo
}

func NewGroupDeviceIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupDeviceIndexLogic {
	return &GroupDeviceIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		GdDB:   relationDB.NewGroupDeviceRepo(ctx),
		DiDB:   relationDB.NewDeviceInfoRepo(ctx),
	}
}

// 获取分组设备信息列表
func (l *GroupDeviceIndexLogic) GroupDeviceIndex(in *dm.GroupDeviceIndexReq) (*dm.GroupDeviceIndexResp, error) {

	var list []*dm.DeviceInfo
	f := relationDB.GroupDeviceFilter{
		GroupID:    in.GroupID,
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
	}
	gd, err := l.GdDB.FindByFilter(l.ctx, f, logic.ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	total, err := l.GdDB.CountByFilter(l.ctx, f)
	if err != nil {
		return nil, err
	}
	for _, v := range gd {
		di, err := l.DiDB.FindOneByFilter(l.ctx, relationDB.DeviceFilter{ProductID: v.ProductID, DeviceNames: []string{v.DeviceName}})
		if err != nil {
			l.Errorf("%s.GroupInfo.DeviceInfoRead failure err=%+v", utils.FuncName(), err)
			continue
		}
		dd := devicemanagelogic.ToDeviceInfo(di)
		list = append(list, dd)
	}

	return &dm.GroupDeviceIndexResp{Total: total, List: list}, nil
}
