package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

func (a *Api) getGroupContact(c echo.Context) (err error) {
	nickname := c.QueryParam("nickname")
	if nickname == "" {
		return errors.New(errors.CodeInvalidParam, "请输入群聊名称")
	}
	result, err := a.service.GetGroupContactByNickname(c.Request().Context(), nickname)
	if err != nil {
		return
	}
	return OK(c, result)
}

func (a *Api) getGroupContactList(c echo.Context) (err error) {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return errors.New(errors.CodeInvalidParam, "offset必须为数字且大于0")
	}
	nickname := c.QueryParam("nickname")

	result, totle, err := a.service.GetGroupContactList(c.Request().Context(), offset, nickname)
	if err != nil {
		return
	}

	return OK(c, map[string]interface{}{
		"list":  result,
		"total": totle,
	})
}

type ReqDelGroup struct {
	Usrname string `json:"user_name" validate:"required"`
}

func (a *Api) delGroupContact(c echo.Context) (err error) {
	var req ReqDelGroup
	if err = c.Bind(&req); err != nil {
		return
	}
	if err = c.Validate(&req); err != nil {
		return
	}

	err = a.service.DelGroupContact(c.Request().Context(), req.Usrname)
	if err != nil {
		return
	}

	return OK(c, "")
}
