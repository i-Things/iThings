package main

import (
	"context"
	"gitee.com/i-Things/share/utils"
	"gitee.com/i-Things/things/service/dgsvr/dgdirect"
)

func main() {
	defer utils.Recover(context.Background())
	svcCtx := dgdirect.GetSvcCtx()
	dgdirect.Run(svcCtx)
}
