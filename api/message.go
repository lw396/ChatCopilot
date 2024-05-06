package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

type ReqSaveMessage struct {
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

	message, err := a.service.ScanMessage(c.Request().Context(), req.UserName)
	if err != nil {
		return
	}

	group, err := a.service.GetGroupContactByUsrname(c.Request().Context(), req.UserName)
	if err != nil {
		return
	}

	group.DBName = message.DBName
	err = a.service.SaveMessageContent(c.Request().Context(), group)
	if err != nil {
		return
	}
	return Created(c, "")
}

func (a *Api) getMessageContentList(c echo.Context) (err error) {
	usrName := c.QueryParam("user_name")
	if usrName == "" {
		return errors.New(errors.CodeInvalidParam, "user_name为空")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil && offset > 0 {
		return errors.New(errors.CodeInvalidParam, "offset必须为数字且大于0")
	}

	result, err := a.service.GetMessageContent(c.Request().Context(), usrName, offset)
	if err != nil {
		return
	}
	return OK(c, result)
}
