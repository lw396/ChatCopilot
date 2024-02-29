package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

func (a *Api) getGroupContact(c echo.Context) (err error) {
	nickname := c.QueryParam("nickname")
	if nickname == "" {
		err = errors.New(errors.CodeInvalidParam, "请输入群聊名称")
		return
	}

	result, err := a.service.GetGroupContact(c.Request().Context(), nickname)
	if err != nil {
		return
	}

	return OK(c, result)
}
