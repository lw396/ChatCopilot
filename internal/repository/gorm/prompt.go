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

func (r *gormRepository) AddPromptCuration(ctx context.Context, req *PromptCuration) (err error) {
	err = r.db.WithContext(ctx).Create(req).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) DelPromptCuration(ctx context.Context, id uint64) (err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).Delete(&PromptCuration{}).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) UpdatePromptCuration(ctx context.Context, req *PromptCuration) (err error) {
	err = r.db.WithContext(ctx).Where("id = ?", req.ID).Updates(req).Error
	if err != nil {
		return
	}
	return
}
