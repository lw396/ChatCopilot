package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

func (a *Api) getMessageInfo(c echo.Context) (err error) {
	userName := c.QueryParam("user_name")
	if userName == "" {
		return errors.New(errors.CodeInvalidParam, "请输入用户名称")
	}
	result, err := a.service.ScanMessage(c.Request().Context(), userName)
	if err != nil {
		return
	}
	return OK(c, result)
}

type ReqSaveMessage struct {
	DBName   string `json:"db_name" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
}

func (a *Api) saveMessageContent(c echo.Context) (err error) {
	var req ReqSaveMessage
	if err = c.Bind(&req); err != nil {
		return
	}
	if err = c.Validate(&req); err != nil {
		return
	}
	err = a.service.SaveMessageContent(c.Request().Context(), req.DBName, req.UserName)
	if err != nil {
		return
	}
	return NoContent(c)
}
