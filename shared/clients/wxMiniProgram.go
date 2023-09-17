package clients

import (
	"context"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	zeroCache "github.com/zeromicro/go-zero/core/stores/cache"
)

type WxMiniProgram struct {
	Open      bool //如果开启则需要初始化为true
	AppID     string
	AppSecret string
}

type MiniProgram = miniprogram.MiniProgram

func NewWxMiniProgram(ctx context.Context, conf WxMiniProgram, redisConf zeroCache.ClusterConf) *MiniProgram {
	if conf.Open == false {
		return nil
	}
	wc := wechat.NewWechat()
	memory := cache.NewRedis(ctx, &cache.RedisOpts{
		Host:     redisConf[0].Host,
		Password: redisConf[0].Pass,
	})
	cfg := &miniConfig.Config{
		AppID:     conf.AppID,
		AppSecret: conf.AppSecret,
		Cache:     memory,
	}
	miniprogram := wc.GetMiniProgram(cfg)
	return miniprogram
}
