package gorm

import (
	"context"
)

func (r *gormRepository) SaveGroupContact(ctx context.Context, content *GroupContact) (err error) {
	err = r.db.WithContext(ctx).Save(content).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetGroupContactByUsrName(ctx context.Context, usrName string) (content *GroupContact, err error) {
	content = &GroupContact{}
	err = r.db.WithContext(ctx).Where("usr_name = ?", usrName).First(content).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetGroupContacts(ctx context.Context) (content []*GroupContact, err error) {
	content = []*GroupContact{}
	err = r.db.WithContext(ctx).Find(&content).Error
	if err != nil {
		return
	}
	return
}
