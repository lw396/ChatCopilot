package sqlite

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/repository"
	"gorm.io/gorm"
)

func (s *SQLite) FindMessage(ctx context.Context, tx *gorm.DB, tableName string) (
	sequence *repository.SQLiteSequence, err error,
) {
	sequence = &repository.SQLiteSequence{}
	if err = tx.First(&sequence, "name = ?", tableName).Error; err != nil {
		return
	}

	return
}

func (s *SQLite) BindMessage(ctx context.Context, tx *gorm.DB, dbName, msgName string) (err error) {
	if s.db[dbName] == nil {
		s.db[dbName] = &DB{tx: tx, msgName: []string{msgName}}
		return
	}
	for _, t := range s.db[dbName].msgName {
		if t == msgName {
			err = errors.New(errors.CodeAuthMessageFound, "message already bind")
			return
		}
	}

	s.db[dbName].msgName = append(s.db[dbName].msgName, msgName)
	return
}

func (s *SQLite) UnbindMessage(ctx context.Context, dbName, msgName string) (err error) {
	db := s.db[dbName]
	if len(db.msgName) == 1 {
		if db.msgName[0] != msgName {
			err = errors.New(errors.CodeAuthMessageFound, "message not bind")
			return
		}

		delete(s.db, dbName)
		return
	}

	var isBind bool
	for i, name := range s.db[dbName].msgName {
		if name != msgName {
			continue
		}
		s.db[dbName].msgName = append(db.msgName[:i], db.msgName[i+1:]...)
		isBind = true
		break
	}

	if !isBind {
		err = errors.New(errors.CodeAuthMessageNotFound, "message not bind")
		return
	}
	return
}