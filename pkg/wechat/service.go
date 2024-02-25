package wechat

type WeChat struct {
	key  string
	path string
}

func NewWeChatClient(key, path string) *WeChat {
	return &WeChat{
		key:  key,
		path: path,
	}
}
