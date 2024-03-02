package gorm

import (
	"context"
)

func (r *gormRepository) CreateMessageContentTable(ctx context.Context, msgName string) (err error) {
	err = r.db.Table(msgName).AutoMigrate(&MessageContent{})
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) SaveMessageContent(ctx context.Context, msgName string, content []*MessageContent) (err error) {
	err = r.db.WithContext(ctx).Table(msgName).Save(&content).Error
	if err != nil {
		return
	}
	return
}
