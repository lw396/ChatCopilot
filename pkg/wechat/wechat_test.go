package wechat

import (
	"context"
	"fmt"
	"testing"

	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func TestGetOpenId(t *testing.T) {
	wc, err := NewOfficialAccount(offConfig.Config{
		AppID:          "wxd64823700a2f7d1f",
		AppSecret:      "3c8fcc8b667dddb04a8cb99e5b47ef96",
		Token:          "",
		EncodingAESKey: "",
		Cache: cache.NewRedis(context.Background(), &cache.RedisOpts{
			Host:        "127.0.0.1",
			Password:    "secret",
			Database:    0,
			MaxIdle:     100,
			MaxActive:   10,
			IdleTimeout: 20,
		}),
	})
	if err != nil {
		t.Error(err)
	}

	openid, err := wc.GetOpenId("0412Mbll22RIAc4Fg0ml2suule42Mblf")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(openid)
}
