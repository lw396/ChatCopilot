package api

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
)

func (api *Api) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		request := c.Request()
		ctx := request.Context()
		authHeader := request.Header.Get("Authorization")
		authScheme := "Bearer "

		if !strings.HasPrefix(authHeader, authScheme) {
			return errors.New(errors.CodeAuthTokenNotFound, "invalid authorization header")
		}

		_, err = api.service.ParseToken(ctx, strings.TrimPrefix(authHeader, authScheme))
		if err != nil {
			return err
		}

		return next(c)
	}
}

type ReqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *Api) Login(c echo.Context) (err error) {
	var req ReqLogin
	if err = c.Bind(&req); err != nil {
		return
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if err = api.service.AuthenticateAccount(req.Username, req.Password); err != nil {
		return
	}

	token, err := api.service.CreateToken(ctx, req.Username)
	if err != nil {
		return
	}

	return Created(c, token)
}
