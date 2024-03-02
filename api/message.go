package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/service"
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
	DBName          string `json:"db_name" validate:"required"`
	UserName        string `json:"user_name" validate:"required"`
	Nickname        string `json:"nickname"`
	HeadImgUrl      string `json:"head_img_url"`
	ChatRoomMemList string `json:"member_list"`
}

func (a *Api) saveMessageContent(c echo.Context) (err error) {
	var req ReqSaveMessage
	if err = c.Bind(&req); err != nil {
		return
	}
	if err = c.Validate(&req); err != nil {
		return
	}
	err = a.service.SaveMessageContent(c.Request().Context(), &service.GroupContact{
		DBName:          req.DBName,
		UsrName:         req.UserName,
		Nickname:        req.Nickname,
		HeadImgUrl:      req.HeadImgUrl,
		ChatRoomMemList: req.ChatRoomMemList,
	})
	if err != nil {
		return
	}
	return Created(c, "")
}
