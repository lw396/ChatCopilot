package gorm

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/model"
)

func (r *gormRepository) GetCopilotConfigByStatus(ctx context.Context, status model.CopilotConfigStatus) (result *CopilotConfig, err error) {
	err = r.db.WithContext(ctx).Where("status = ?", status).First(&result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) AddChatCopilot(ctx context.Context, copilot *ChatCopilot) (err error) {
	err = r.db.WithContext(ctx).Omit("Prompt").Create(copilot).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetChatCopilotList(ctx context.Context) (result []*ChatCopilot, err error) {
	err = r.db.WithContext(ctx).Preload("Prompt").Find(result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetChatCopilot(ctx context.Context, id int64) (result *ChatCopilot, err error) {
	err = r.db.WithContext(ctx).Preload("Prompt").Where("id = ?", id).First(result).Error
	if err != nil {
		return
	}
	return
}
