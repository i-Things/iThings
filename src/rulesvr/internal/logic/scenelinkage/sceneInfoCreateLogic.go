package scenelinkagelogic

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/src/rulesvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/rulesvr/internal/svc"
	"github.com/i-Things/things/src/rulesvr/pb/rule"
	"github.com/zeromicro/go-zero/core/logx"
)

type SceneInfoCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	SiDB *relationDB.SceneInfoRepo
}

func NewSceneInfoCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SceneInfoCreateLogic {
	return &SceneInfoCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		SiDB:   relationDB.NewSceneInfoRepo(ctx),
	}
}

func (l *SceneInfoCreateLogic) SceneInfoCreate(in *rule.SceneInfo) (*rule.WithID, error) {
	do, err := ToSceneDo(in)
	if err != nil {
		return nil, err
	}
	_, err = l.SiDB.FindOneByName(l.ctx, do.Name)
	if err == nil {
		return nil, errors.Parameter.AddMsg("场景名字重复")
	}
	if !errors.Cmp(err, errors.NotFind) {
		return nil, err
	}
	err = do.Validate()
	if err != nil {
		return nil, err
	}
	id, err := l.SiDB.Insert(l.ctx, do)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.SceneTimerControl.Create(do)
	if err != nil {
		return nil, err
	}
	return &rule.WithID{Id: id}, err
}
