package sqlite

import (
	"context"
)

func (s *SQLite) GetGroupContactByNickname(ctx context.Context, nickname string) (
	result []*GroupContact, err error) {
	err = s.db[GroupDB].Where("nickname LIKE ?", "%"+nickname+"%").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
