package interact

import (
	"context"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/things/service/apisvr/internal/svc"
	"gitee.com/i-Things/things/service/apisvr/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendPropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendPropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPropertyLogic {
	return &SendPropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctxs.WithDefaultRoot(ctx),
		svcCtx: svcCtx,
	}
}

func (l *SendPropertyLogic) SendProperty(req *types.DeviceInteractSendPropertyReq) (resp *types.DeviceInteractSendPropertyResp, err error) {
	return NewPropertyControlSendLogic(l.ctx, l.svcCtx).PropertyControlSend(req)
}
