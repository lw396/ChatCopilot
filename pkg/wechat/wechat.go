package wechat

import (
	"context"
	"fmt"
)

type WeChatClient interface {
	GetWechatConfig(ctx context.Context) (err error)
}

func (wc *WeChat) GetWechatConfig(ctx context.Context) (err error) {
	fmt.Println(wc)
	return
}
