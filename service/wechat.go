package service

import "context"

func (a *Service) GetConfig(ctx context.Context) (err error) {
	err = a.wechat.GetWechatConfig(ctx)
	return
}
