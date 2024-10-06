package auth5

import (
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/result"
	"gitee.com/i-Things/share/utils"
	"gitee.com/i-Things/things/service/apisvr/internal/logic/things/device/auth5"
	"net/http"

	"gitee.com/i-Things/things/service/apisvr/internal/svc"
	"gitee.com/i-Things/things/service/apisvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceAuth5LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Http(w, r, nil, errors.Parameter.AddMsg(err.Error()))
			return
		}

		l := auth5.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		l.Infof("%s req=%v resp=%v err=%v", utils.FuncName(), utils.Fmt(req), utils.Fmt(resp), err)
		result.HttpWithoutWrap(w, r, resp, err)
	}
}
