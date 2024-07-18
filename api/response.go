package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/lw396/ChatCopilot/internal/errors"
)

func Success(c echo.Context, statusCode int, data interface{}, metaKVs ...interface{}) error {
	return c.JSON(statusCode, success(data, metaKVs...))
}

func OK(c echo.Context, data interface{}, metaKVs ...interface{}) error {
	return Success(c, http.StatusOK, data, metaKVs...)
}

func Created(c echo.Context, data interface{}, metaKVs ...interface{}) error {
	return Success(c, http.StatusCreated, data, metaKVs...)
}

func Paginate(c echo.Context, data interface{}, total int64, metaKVs ...interface{}) error {
	metaKVs = append([]interface{}{"total", total}, metaKVs...)
	return Success(c, http.StatusOK, data, metaKVs...)
}

func NoContent(c echo.Context) error {
	return Success(c, http.StatusNoContent, nil)
}

func StreamResponse(c echo.Context, ch chan interface{}) error {
	c.Response().Header().Set(echo.HeaderContentType, "application/x-ndjson")
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())
	for val := range ch {
		if err := enc.Encode(val); err != nil {
			return err
		}
		c.Response().Flush()
	}
	return nil
}

func success(data interface{}, metaKVs ...interface{}) Response {
	resp := Response{
		Code: 0,
		Data: data,
		Meta: make(map[string]interface{}),
	}

	l := len(metaKVs)
	for i := 0; i < l; i += 2 {
		if key, ok := metaKVs[i].(string); ok && i+1 < l {
			resp.Meta[key] = metaKVs[i+1]
		}
	}
	return resp
}

func fail(code int, msg string) Response {
	resp := Response{
		Code:    code,
		Meta:    make(map[string]interface{}),
		Message: msg,
	}
	return resp
}

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Data    interface{}            `json:"data"`
	Meta    map[string]interface{} `json:"meta"`
}

func HTTPErrorHandler(e error, ctx echo.Context) {
	code := errors.CodeGeneral
	statusCode := http.StatusBadRequest
	msg := e.Error()

	if he, ok := e.(*echo.HTTPError); ok {
		code = errors.CodeInvalidParam
		msg = fmt.Sprintf("%v", he.Message)
		statusCode = he.Code
	}

	if mc, ok := e.(*errors.Error); ok {
		code = mc.Code()
		statusCode = mc.HTTPStatusCode()
	}

	if _, ok := e.(*mysql.MySQLError); ok {
		code = errors.CodeDB
		msg = "internal error"
	}

	_ = ctx.JSON(statusCode, fail(code, msg))
}
