package api

import (
	"github.com/labstack/echo/v4"
)

func (api *Api) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		return next(c)
	}
}
