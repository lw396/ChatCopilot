package service

import (
	"context"

	db "github.com/lw396/ChatCopilot/internal/repository/gorm"
)

func (a *Service) AddPromptCuration(ctx context.Context, req *db.PromptCuration) (err error) {
	if err = a.rep.AddPromptCuration(ctx, req); err != nil {
		return
	}
	return
}

type PromptCuration struct {
	Id     uint64 `json:"id"`
	Title  string `json:"title"`
	Prompt string `json:"prompt"`
	Start  uint8  `json:"start"`
}

func (a *Service) GetPromptCurationList(ctx context.Context, offset, limit int) (result []*PromptCuration, total int64, err error) {
	prompt, total, err := a.rep.GetPromptCurationList(ctx, offset, limit)
	if err != nil {
		return
	}

	result = make([]*PromptCuration, len(prompt))
	for i, v := range prompt {
		result[i] = &PromptCuration{
			Id:     v.ID,
			Prompt: v.Prompt,
			Start:  v.Start,
			Title:  v.Title,
		}
	}
	return
}

func (a *Service) DelPromptCuration(ctx context.Context, id uint64) (err error) {
	if err = a.rep.DelPromptCuration(ctx, id); err != nil {
		return
	}
	return
}

func (a *Service) UpdatePromptCuration(ctx context.Context, req *db.PromptCuration) (err error) {
	if err = a.rep.UpdatePromptCuration(ctx, req); err != nil {
		return
	}
	return
}
