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

func (a *Api) getMessageImage(c echo.Context) (err error) {
	path := c.QueryParam("path")
	if path == "" {
		return errors.New(errors.CodeInvalidParam, "path为空")
	}

	image, err := a.service.GetMessageImage(c.Request().Context(), path)
	if err != nil {
		return
	}
	return c.File(image)
}

func (a *Api) getMessageSticker(c echo.Context) (err error) {
	path, url := c.QueryParam("path"), c.QueryParam("url")
	if path == "" || url == "" {
		return errors.New(errors.CodeInvalidParam, "参数错误")
	}

	image, err := a.service.GetMessageSticker(c.Request().Context(), path, url)
	if err != nil {
		return
	}

	return c.File(image)
}
