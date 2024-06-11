package gorm

import (
	"context"
)

func (r *gormRepository) CreateMessageContentTable(ctx context.Context, msgName string) (err error) {
	err = r.db.Table(msgName).Migrator().CreateTable(&MessageContent{})
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

func (r *gormRepository) GetNewMessageContent(ctx context.Context, msgName string) (result *MessageContent, err error) {
	result = &MessageContent{}
	err = r.db.WithContext(ctx).Table(msgName).Order("local_id desc").First(result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) UpdateMessageContent(ctx context.Context, msgName string, content *MessageContent) (err error) {
	err = r.db.WithContext(ctx).Table(msgName).Where("local_id = ?", content.LocalID).
		Update("translate", content.Translate).Error
	return
}

func (r *gormRepository) GetMessageContentList(ctx context.Context, msgName string, offset int) (result []*MessageContent, err error) {
	err = r.db.WithContext(ctx).Table(msgName).Order("local_id desc").Limit(30).Offset(offset).
		Find(&result).Error
	if err != nil {
		return
	}
	return
}

func (r *gormRepository) DelMessageContentTable(ctx context.Context, msgName string) (err error) {
	return r.db.Exec("DROP TABLE " + msgName).Error
}
