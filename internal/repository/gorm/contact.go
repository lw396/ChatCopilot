package gorm

import (
	"context"
)

func (r *gormRepository) SaveContactPerson(ctx context.Context, contact *ContactPerson) (err error) {
	err = r.db.WithContext(ctx).Save(contact).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetContactPersonByUsrName(ctx context.Context, usrName string) (result *ContactPerson, err error) {
	result = &ContactPerson{}
	err = r.db.WithContext(ctx).Where("usr_name = ?", usrName).First(result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) GetContactPersons(ctx context.Context, nickname string, offset int) (result []*ContactPerson, total int64, err error) {
	result = []*ContactPerson{}
	tx := r.db.WithContext(ctx).Model(&ContactPerson{})
	if offset >= 0 {
		tx = tx.Limit(10).Offset(offset)
	}
	tx = tx.Count(&total)

	if nickname != "" {
		tx = tx.Where("nickname LIKE ?", "%"+nickname+"%")
	}
	err = tx.Count(&total).Find(&result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) DelContactPersonByUsrName(ctx context.Context, usrame string) (err error) {
	return r.db.WithContext(ctx).Where("usr_name = ?", usrame).Unscoped().Delete(&ContactPerson{}).Error
}
