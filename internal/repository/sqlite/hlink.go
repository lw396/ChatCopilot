package sqlite

import "context"

func (s *SQLite) GetHinkMediaByMediaMd5(ctx context.Context, mediaMd5 string) (result *HlinkMediaRecord, err error) {
	err = s.db[HlinkDB].Where("mediaMd5 = ?", mediaMd5).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
