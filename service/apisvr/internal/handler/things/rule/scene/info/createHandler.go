package info

import (
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/result"
	"gitee.com/i-Things/things/service/apisvr/internal/logic/things/rule/scene/info"
	"gitee.com/i-Things/things/service/apisvr/internal/svc"
	"gitee.com/i-Things/things/service/apisvr/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SceneInfoCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.Http(w, r, nil, errors.Parameter.WithMsg("入参不正确:"+err.Error()))
			return
		}

		l := info.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		result.Http(w, r, resp, err)
	}
}
