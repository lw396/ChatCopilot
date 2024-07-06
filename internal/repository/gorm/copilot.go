package gorm

import (
	"context"
)

func (r *gormRepository) GetCopilotConfigByModel(ctx context.Context, model string) (result *CopilotConfig, err error) {
	err = r.db.WithContext(ctx).Where("model = ?", model).First(&result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) AddChatCopilot(ctx context.Context, copilot *ChatCopilot) (err error) {
	err = r.db.WithContext(ctx).Create(copilot).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetChatCopilotList(ctx context.Context) (result []*ChatCopilot, err error) {
	err = r.db.WithContext(ctx).Find(result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetChatCopilot(ctx context.Context, id int64) (result *ChatCopilot, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).First(result).Error
	if err != nil {
		return
	}
	return
}
