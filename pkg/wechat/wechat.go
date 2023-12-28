package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

type WechatClient interface {
	GetOpenId(code string) (openid string, err error)
	SendMessage(msg *message.TemplateMessage) (err error)
}

type ExampleOfficialAccount struct {
	wc              *wechat.Wechat
	officialAccount *officialaccount.OfficialAccount
}

func NewOfficialAccount(cfg config.Config) (WechatClient, error) {
	stableHandle := credential.NewStableAccessToken(
		cfg.AppID,
		cfg.AppSecret,
		credential.CacheKeyOfficialAccountPrefix,
		cfg.Cache,
	)

	wc := wechat.NewWechat()
	officialAccount := wc.GetOfficialAccount(&cfg)
	officialAccount.SetAccessTokenHandle(stableHandle)

	return &ExampleOfficialAccount{
		wc:              wc,
		officialAccount: officialAccount,
	}, nil
}
