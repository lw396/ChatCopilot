package sqlite

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository"
)

func (s *SQLite) GetGroupContactByNickname(ctx context.Context, nickname string) (
	result []*repository.GroupContact,
	err error,
) {
	result = []*repository.GroupContact{}
	err = s.db[GroupDB].tx.Find(result, "nickname LIKE ?", "%"+nickname+"%").Error
	if err != nil {
		return
	}

	return
}
