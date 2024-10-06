package info

import (
	"gitee.com/i-Things/share/utils"
	"gitee.com/i-Things/things/service/apisvr/internal/types"
	"gitee.com/i-Things/things/service/udsvr/pb/ud"
)

func ToSceneTypes(in *ud.SceneInfo) *types.SceneInfo {
	return utils.Copy[types.SceneInfo](in)
}

func ToScenePb(in *types.SceneInfo) *ud.SceneInfo {
	return utils.Copy[ud.SceneInfo](in)
}
