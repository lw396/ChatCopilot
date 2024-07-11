package service

import (
	"context"

	db "github.com/lw396/WeComCopilot/internal/repository/gorm"
)

func (a *Service) AddPromptCuration(ctx context.Context, req *db.PromptCuration) (err error) {
	if err = a.rep.AddPromptCuration(ctx, req); err != nil {
		return
	}
	return
}

func (a *Service) GetPromptCurationList(ctx context.Context, offset, limit int) (prompt []*db.PromptCuration, err error) {
	prompt, err = a.rep.GetPromptCurationList(ctx, offset, limit)
	if err != nil {
		return
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
