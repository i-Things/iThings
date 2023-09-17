package user

import (
	"context"
	"github.com/i-Things/things/shared/ctxs"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UserInfo) error {
	//超管可以修改其他用户的角色，并且任何用户不能修改本身的角色
	roleID := int64(0)
	userCtx := ctxs.GetUserCtx(l.ctx)
	if !userCtx.IsOpen && userCtx.Role == int64(def.RoleIDSuper) && userCtx.UserID != req.UserID {
		roleID = req.Role
	}
	_, err := l.svcCtx.UserRpc.UserUpdate(l.ctx, &sys.UserInfo{
		UserID:     req.UserID,
		UserName:   req.UserName,
		Email:      req.Email,
		NickName:   req.NickName,
		City:       req.City,
		Country:    req.Country,
		Province:   req.Province,
		Language:   req.Language,
		HeadImgUrl: req.HeadImgUrl,
		Role:       roleID,
		Sex:        req.Sex,
		IsAllData:  req.IsAllData,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.user.upadte failure err=%+v", utils.FuncName(), er)
		return er
	}
	return nil
}
