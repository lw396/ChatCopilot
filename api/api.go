package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lw396/ChatCopilot/service"
)

type Api struct {
	port    int64
	service *service.Service
}

type Config struct {
	App  *service.Service
	Port int64
}

func New(c Config) *Api {
	return &Api{
		port:    c.Port,
		service: c.App,
	}
}

func (api *Api) Run() error {
	engine := echo.New()
	engine.HTTPErrorHandler = HTTPErrorHandler
	engine.Validator = NewValidator()

	engine.Use(middleware.CORS())
	engine.Use(middleware.Recover())
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderAuthorization},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST},
	}))

	engine.POST("/auth/login", api.login)

	v1 := engine.Group("/v1", api.authenticate)
	{
		v1.GET("/user", api.getUser)
		// 群
		// 获取群聊名称列表
		v1.GET("/group_contact", api.getGroupContact)
		// 保存群聊聊天记录
		v1.POST("/group_contact", api.saveGroupContact)
		// 删除群聊信息及记录
		v1.DELETE("/group_contact", api.delGroupContact)
		// 查看群聊列表
		v1.GET("/group_contact_list", api.getGroupContactList)

		// 联系人
		// 获取联系人信息列表
		v1.GET("/contact_person", api.getContactPerson)
		// 保存联系人聊天记录
		v1.POST("/contact_person", api.saveContactPerson)
		// 删除联系人信息及记录
		v1.DELETE("/contact_person", api.delContactPerson)
		// 查看联系人列表
		v1.GET("/contact_person_list", api.getContactPersonList)

		// 聊天记录
		// 查看聊天记录列表
		v1.GET("/message_content_list", api.getMessageContentList)
		// 查看图片
		v1.GET("/message_image", api.getMessageImage)
		// 查看表情包
		v1.GET("/message_sticker", api.getMessageSticker)
		// 播放语音
		v1.GET("/message_voice", api.getMessageVoice)

		// 聊天助手
		// 添加聊天助手
		v1.POST("/chat_copilot", api.addChatCopilot)
		// 获取聊天提示
		v1.POST("/chat_tips", api.getChatTips)

		// 提示词
		// 新增提示词
		v1.POST("/prompt_curation", api.addPromptCuration)
		// 删除提示词
		v1.DELETE("/prompt_curation", api.delPromptCuration)
		// 修改提示词
		v1.PUT("/prompt_curation", api.updatePromptCuration)
		// 获取提示词列表
		v1.GET("/prompt_curation_list", api.getPromptCurationList)
	}

	return engine.Start(fmt.Sprintf(":%d", api.port))
}
