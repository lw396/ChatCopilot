package wechat

import (
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func (wx *ExampleOfficialAccount) GetOpenId(code string) (openid string, err error) {
	accountAuth := wx.officialAccount.GetOauth()
	result, err := accountAuth.GetUserAccessToken(code)
	if err != nil {
		return
	}

	openid = result.OpenID
	return
}

func (wx *ExampleOfficialAccount) SendMessage(msg *message.TemplateMessage) (err error) {
	template := wx.officialAccount.GetTemplate()
	if _, err = template.Send(msg); err != nil {
		return err
	}

	return
}
