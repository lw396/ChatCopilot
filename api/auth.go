package api

import (
	"github.com/labstack/echo/v4"
)

// 验证 PHPSESSID
func (api *Api) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		return next(c)
	}
}
