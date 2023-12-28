package gorm

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/db"
)

func (r *gormRepository) classifyHelper() *db.Helper[repository.Classify] {
	return db.NewHelper[repository.Classify](r.db)
}

func (r *gormRepository) GetClassify(ctx context.Context, types uint8) ([]*repository.Classify, error) {
	return r.classifyHelper().Where("is_del = 1 AND top = 1 AND types = ?", types).Order("weigh").Find(ctx)
}

func (r *gormRepository) setMealHelper() *db.Helper[repository.SetMeal] {
	return db.NewHelper[repository.SetMeal](r.db)
}

func (r *gormRepository) GetSetMeal(ctx context.Context, types uint8) ([]*repository.SetMeal, error) {
	return r.setMealHelper().Where("is_del = 1 AND top = 1 AND type = ? ", types).Order("weigh").Find(ctx)
}
