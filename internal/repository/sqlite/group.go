package sqlite

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/db"
)

func (s *SQLite) groupContactHelper() *db.Helper[repository.GroupContact] {
	return db.NewHelper[repository.GroupContact](s.db[GroupDB].tx)
}

func (s *SQLite) GetGroupContactByNickname(ctx context.Context, nickname string) (
	result []*repository.GroupContact, err error) {
	return s.groupContactHelper().Where("nickname LIKE ?", "%"+nickname+"%").Find(ctx)
}
