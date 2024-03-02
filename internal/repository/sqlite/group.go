package sqlite

import (
	"context"

	"github.com/lw396/WeComCopilot/pkg/db"
)

func (s *SQLite) groupContactHelper() *db.Helper[GroupContact] {
	return db.NewHelper[GroupContact](s.db[GroupDB].tx)
}

func (s *SQLite) GetGroupContactByNickname(ctx context.Context, nickname string) (
	[]*GroupContact, error) {
	return s.groupContactHelper().Where("nickname LIKE ?", "%"+nickname+"%").Find(ctx)
}
