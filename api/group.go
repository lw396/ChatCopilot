package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

// getGroupContact	godoc
//
//	@Summary		获取群聊名称列表
//	@Description	获取群聊名称列表通过输入群聊名称列表
//	@Tags			group
//	@Accept			json
//	@Produce		json
func (a *Api) getGroupContact(c echo.Context) (err error) {
	nickname := c.QueryParam("nickname")
	if nickname == "" {
		return errors.New(errors.CodeInvalidParam, "请输入群聊名称")
	}
	result, err := a.service.GetGroupContact(c.Request().Context(), nickname)
	if err != nil {
		return
	}
	return OK(c, result)
}
