package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

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
