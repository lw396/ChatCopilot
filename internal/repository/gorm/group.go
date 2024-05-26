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

func (r *gormRepository) GetGroupContacts(ctx context.Context, nickname string, offset int) (content []*GroupContact, total int64, err error) {
	content = []*GroupContact{}
	tx := r.db.WithContext(ctx).Model(&GroupContact{}).Count(&total)
	if offset > 0 {
		tx = tx.Limit(10).Offset(offset)
	}
	if nickname != "" {
		tx = tx.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	err = tx.Find(&content).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) DelGroupContactByUsrName(ctx context.Context, usrName string) (err error) {
	return r.db.WithContext(ctx).Where("usr_name = ?", usrName).Unscoped().Delete(&GroupContact{}).Error
}
