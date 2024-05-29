package sqlite

import (
	"context"

	"gorm.io/gorm"
)

func (s *SQLite) BindMessageDB(ctx context.Context, tx *gorm.DB, dbName string) (err error) {
	if s.db[dbName] != nil {
		return
	}
	s.db[dbName] = tx
	return
}

func (s *SQLite) UnbindMessageDB(ctx context.Context, dbName string) {
	delete(s.db, dbName)
}

func (s *SQLite) CheckMessageExistDB(ctx context.Context, tx *gorm.DB, userName string) (result *SQLiteSequence, err error) {
	result = &SQLiteSequence{}
	err = tx.WithContext(ctx).Where("name = ?", userName).First(result).Error
	if err != nil {
		if err.Error() == "no such table: sqlite_sequence" {
			err = gorm.ErrRecordNotFound
		}
		return
	}
	return
}

func (s *SQLite) GetMessageContent(ctx context.Context, dbName, msgName string) (result []*MessageContent, err error) {
	err = s.db[dbName].WithContext(ctx).Table(msgName).Order("mesLocalID").Find(&result).Error
	if err != nil {
		return
	}
	return
}

func (s *SQLite) GetUnsyncMessageContent(ctx context.Context, dbName, msgName string, newId int64) (result []*MessageContent, err error) {
	err = s.db[dbName].WithContext(ctx).Table(msgName).Order("mesLocalID").Where("mesLocalID > ?", newId).Find(&result).Error
	if err != nil {
		return
	}
	return
}
