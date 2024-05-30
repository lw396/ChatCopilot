package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

func (a *Api) getContactPerson(c echo.Context) (err error) {
	nickname := c.QueryParam("nickname")
	if nickname == "" {
		return errors.New(errors.CodeInvalidParam, "请输入联系人昵称")
	}

	result, err := a.service.GetContactPersonByNickname(c.Request().Context(), nickname)
	if err != nil {
		return
	}
	return OK(c, result)
}

type ReqSaveContact struct {
	Usrname string `json:"user_name" validate:"required"`
}

func (a *Api) saveContactPerson(c echo.Context) (err error) {
	var req ReqSaveContact
	if err = c.Bind(&req); err != nil {
		return
	}
	if err = c.Validate(&req); err != nil {
		return
	}

	message, err := a.service.ScanMessage(c.Request().Context(), req.Usrname)
	if err != nil {
		return
	}

	contact, err := a.service.GetContactPersonByUsrname(c.Request().Context(), req.Usrname)
	if err != nil {
		return
	}

	contact.DBName = message.DBName
	err = a.service.SaveContactPerson(c.Request().Context(), contact)
	if err != nil {
		return
	}

	return Created(c, "")
}

type ReqDelContact struct {
	Usrname string `json:"user_name" validate:"required"`
}

func (a *Api) delContactPerson(c echo.Context) (err error) {
	var req ReqDelContact
	if err = c.Bind(&req); err != nil {
		return
	}
	if err = c.Validate(&req); err != nil {
		return
	}

	err = a.service.DelContactPerson(c.Request().Context(), req.Usrname)
	if err != nil {
		return
	}

	return OK(c, "")
}

func (a *Api) getContactPersonList(c echo.Context) (err error) {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return errors.New(errors.CodeInvalidParam, "offset必须为数字且大于0")
	}
	nickname := c.QueryParam("nickname")

	result, totle, err := a.service.GetContactPersonList(c.Request().Context(), offset, nickname)
	if err != nil {
		return
	}

	return Paginate(c, result, totle)
}
