package gorm

import (
	"context"

	"github.com/lw396/WeComCopilot/pkg/db"
)

func (r *gormRepository) messageContentHelper() *db.Helper[MessageContent] {
	return db.NewHelper[MessageContent](r.db)
}
func (r *gormRepository) SaveMessageContent(ctx context.Context, msgName string, content []*MessageContent) (err error) {
	return
}
