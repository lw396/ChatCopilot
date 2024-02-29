package sqlite

import (
	"context"
	"fmt"

	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/sqlcipher"
	"gorm.io/gorm"
)

type SQLiteClient interface {
	OpenDB(ctx context.Context, dbName string) (tx *gorm.DB, err error)
	FindMessage(ctx context.Context, tx *gorm.DB, msgName string) (sequence *repository.SQLiteSequence, err error)
	BindMessage(ctx context.Context, tx *gorm.DB, dbName string, msgName string) (err error)
	UnbindMessage(ctx context.Context, dbName string, msgName string) (err error)
}

func (s *SQLite) OpenDB(ctx context.Context, dbName string) (tx *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s/%s?_pragma_key=x'%s'", s.path, dbName, s.key)

	tx, err = gorm.Open(sqlcipher.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	return
}

func (s *SQLite) FindMessage(ctx context.Context, tx *gorm.DB, tableName string) (
	sequence *repository.SQLiteSequence, err error,
) {
	sequence = &repository.SQLiteSequence{}
	if err = tx.First(&sequence, "name = ?", tableName).Error; err != nil {
		return
	}

	return
}

func (s *SQLite) BindMessage(ctx context.Context, tx *gorm.DB, dbName string, msgName string) (err error) {
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

func (s *SQLite) UnbindMessage(ctx context.Context, dbName string, msgName string) (err error) {
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
