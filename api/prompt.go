package api

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/repository/gorm"
)

type RepPromptCuration struct {
	Id     uint64 `json:"id"`
	Title  string `json:"title" validate:"required"`
	Prompt string `json:"prompt" validate:"required"`
	Start  uint8  `json:"start"`
}

func (a *Api) addPromptCuration(c echo.Context) (err error) {
	req := &RepPromptCuration{}
	if err = c.Bind(req); err != nil {
		return
	}
	if err = c.Validate(req); err != nil {
		return
	}

	if err = a.service.AddPromptCuration(c.Request().Context(), &gorm.PromptCuration{
		Prompt: req.Prompt,
		Start:  req.Start,
		Title:  req.Title,
	}); err != nil {
		return
	}

	return Created(c, "")
}

func (a *Api) getPromptCurationList(c echo.Context) (err error) {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return errors.New(errors.CodeInvalidParam, "offset必须为数字")
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return errors.New(errors.CodeInvalidParam, "limit必须为数字")
	}

	list, err := a.service.GetPromptCurationList(c.Request().Context(), offset, limit)
	if err != nil {
		return
	}

	return OK(c, list)
}

type RepDelPrompt struct {
	Id uint64 `json:"id" validate:"required"`
}

func (a *Api) delPromptCuration(c echo.Context) (err error) {
	req := &RepDelPrompt{}
	if err = c.Bind(req); err != nil {
		return
	}
	if err = c.Validate(req); err != nil {
		return
	}

	if err = a.service.DelPromptCuration(c.Request().Context(), req.Id); err != nil {
		return
	}

	return
}

func (a *Api) updatePromptCuration(c echo.Context) (err error) {
	req := &RepPromptCuration{}
	if err = c.Bind(req); err != nil {
		return
	}
	if err = c.Validate(req); err != nil {
		return
	}

	if err = a.service.UpdatePromptCuration(c.Request().Context(), &gorm.PromptCuration{
		Model:  gorm.Model{ID: req.Id},
		Prompt: req.Prompt,
		Start:  req.Start,
		Title:  req.Title,
	}); err != nil {
		return
	}
	return
}
