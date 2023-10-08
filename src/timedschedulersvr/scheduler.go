package main

import (
	"context"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/timedschedulersvr/timedschedulerdirect"
)

func main() {
	defer utils.Recover(context.Background())
	ctx := timedschedulerdirect.GetSvcCtx()
	timedschedulerdirect.Run(ctx)
}
