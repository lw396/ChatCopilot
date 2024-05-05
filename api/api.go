package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lw396/WeComCopilot/service"
	_ "github.com/swaggo/echo-swagger/example/docs"
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
		// 获取群聊名称列表
		v1.GET("/group_contact", api.getGroupContact)
		// 获取群聊基本信息
		v1.GET("/message_info", api.getMessageInfo)
		// 保存群聊聊天记录
		v1.POST("/message_content", api.saveMessageContent)
		// 查看同步群聊列表
		v1.GET("/group_contact_list", api.getGroupContactList)
		// 删除群聊信息及记录
		v1.DELETE("/group_contact", api.delGroupContact)
		// 查看群聊记录列表
		v1.GET("/message_content_list", api.getMessageContentList)
	}

	return engine.Start(fmt.Sprintf(":%d", api.port))
}
