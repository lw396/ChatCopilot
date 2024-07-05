package gorm

import "context"

func (r *gormRepository) GetPromptCurationList(ctx context.Context, offset, limit int) (result []*PromptCuration, err error) {
	err = r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetPromptCuration(ctx context.Context, id int64) (result *PromptCuration, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return
	}
	return
}
