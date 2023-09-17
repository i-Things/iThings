package devicegrouplogic

import (
	"context"
	"github.com/i-Things/things/src/dmsvr/internal/logic"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupInfoIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	GiDB *relationDB.GroupInfoRepo
}

func NewGroupInfoIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoIndexLogic {
	return &GroupInfoIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		GiDB:   relationDB.NewGroupInfoRepo(ctx),
	}
}

// 获取分组信息列表
func (l *GroupInfoIndexLogic) GroupInfoIndex(in *dm.GroupInfoIndexReq) (*dm.GroupInfoIndexResp, error) {
	f := relationDB.GroupInfoFilter{
		GroupName:   in.GroupName,
		ParentID:    in.ParentID,
		Tags:        in.Tags,
		WithProduct: true,
	}
	ros, err := l.GiDB.FindByFilter(l.ctx, f, logic.ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	total, err := l.GiDB.CountByFilter(l.ctx, f)
	if err != nil {
		return nil, err
	}
	info := make([]*dm.GroupInfo, 0, len(ros))
	for _, ro := range ros {
		info = append(info, ToGroupInfoPb(ro))
	}
	f.ParentID = 0
	rosAll, err := l.GiDB.FindByFilter(l.ctx, f, logic.ToPageInfo(in.Page))
	if err != nil {
		return nil, err
	}
	infoAll := make([]*dm.GroupInfo, 0, len(rosAll))
	for _, ro := range rosAll {
		infoAll = append(infoAll, ToGroupInfoPb(ro))
	}
	return &dm.GroupInfoIndexResp{List: info, Total: total, ListAll: infoAll}, nil
}
