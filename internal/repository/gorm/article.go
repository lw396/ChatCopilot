package gorm

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/db"
)

func (r *gormRepository) articleHelper() *db.Helper[repository.Article] {
	return db.NewHelper[repository.Article](r.db)
}

func (r *gormRepository) GetArticleByTitle(ctx context.Context, title string) (*repository.Article, error) {
	return r.articleHelper().Where("title = ?", title).First(ctx)
}
