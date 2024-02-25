package api

import (
	"github.com/labstack/echo/v4"
)

func (api *Api) Test(c echo.Context) (err error) {
	ctx := c.Request().Context()

	err = api.service.GetConfig(ctx)
	if err != nil {
		return
	}
	return nil
}
