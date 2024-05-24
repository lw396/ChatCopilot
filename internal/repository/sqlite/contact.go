package sqlite

import "context"

func (s *SQLite) GetContactPersonByNickname(ctx context.Context, nickname string) (
	result []*ContactPerson, err error) {
	err = s.db[ContactDB].Where("nickname LIKE ?", "%"+nickname+"%").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return
}

func (s *SQLite) GetContactPersonByUsrname(ctx context.Context, usrname string) (
	result *ContactPerson, err error) {
	err = s.db[ContactDB].Where("usr_name = ?", usrname).First(&result).Error
	if err != nil {
		return nil, err
	}
	return
}
