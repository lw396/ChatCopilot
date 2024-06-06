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

func (r *gormRepository) GetGroupContactByUsrName(ctx context.Context, usrName string) (result *GroupContact, err error) {
	result = &GroupContact{}
	err = r.db.WithContext(ctx).Where("usr_name = ?", usrName).First(result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetGroupContacts(ctx context.Context, nickname string, offset int) (result []*GroupContact, total int64, err error) {
	result = []*GroupContact{}
	tx := r.db.WithContext(ctx).Model(&GroupContact{})
	if nickname != "" {
		tx = tx.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	tx = tx.Count(&total)

	if offset >= 0 {
		tx = tx.Limit(10).Offset(offset)
	}
	err = tx.Find(&result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) DelGroupContactByUsrName(ctx context.Context, usrName string) (err error) {
	return r.db.WithContext(ctx).Where("usr_name = ?", usrName).Unscoped().Delete(&GroupContact{}).Error
}
