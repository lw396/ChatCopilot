package api

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/model"
	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"
)

type ReqAddCopilot struct {
	Usrname  string         `json:"user_name" validate:"required"`
	PromptID int64          `json:"prompt_id" validate:"required"`
	ChatType model.ChatType `json:"chat_type" validate:"required"`
}

func (a *Api) addChatCopilot(c echo.Context) (err error) {
	var req ReqAddCopilot
	if err = c.Bind(&req); err != nil {
		return
	}
	if err = c.Validate(&req); err != nil {
		return
	}

	if err = a.service.AddChatCopilot(c.Request().Context(), &mysql.ChatCopilot{
		UsrName:  req.Usrname,
		Type:     req.ChatType,
		PromptID: req.PromptID,
	}); err != nil {
		return
	}

	return Created(c, "")
}

func (a *Api) getChatTips(c echo.Context) (err error) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	usrname := c.QueryParam("user_name")
	if usrname == "" {
		return errors.New(errors.CodeInvalidParam, "user_name 为空")
	}

	pipeReader, pipeWriter := io.Pipe()

	if err = a.service.GetChatTips(c.Request().Context(), usrname, pipeWriter); err != nil {
		return
	}

	return c.Stream(http.StatusOK, "application/x-ndjson", pipeReader)
}
