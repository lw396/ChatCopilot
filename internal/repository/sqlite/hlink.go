package sqlite

import (
	"context"
)

func (s *SQLite) GetHinkMediaByMediaMd5(ctx context.Context, mediaMd5 string) (result *HlinkMediaRecord, err error) {
	result = &HlinkMediaRecord{}
	err = s.db[HlinkDB].Preload("Detail").Where("mediaMd5 = ?", mediaMd5).First(result).Error
	if err != nil {
		return nil, err
	}
	return
}
