package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lw396/WeComCopilot/service"
	echoSwagger "github.com/swaggo/echo-swagger"
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

//	@title			Chat Copilot API
//	@version		v1
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		petstore.swagger.io
// @BasePath	/v2
func (api *Api) Run() error {
	engine := echo.New()
	engine.HTTPErrorHandler = HTTPErrorHandler
	engine.Validator = NewValidator()

	engine.Use(middleware.CORS())
	engine.Use(middleware.Recover())
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST},
	}))

	v1 := engine.Group("/v1")

	v1.GET("/swagger/*", echoSwagger.WrapHandler)

	v1.GET("/group_contact", api.getGroupContact)
	v1.GET("/message_info", api.getMessageInfo)
	v1.POST("/message_content", api.saveMessageContent)

	return engine.Start(fmt.Sprintf(":%d", api.port))
}
