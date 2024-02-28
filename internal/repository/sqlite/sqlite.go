package sqlite

import (
	"context"
	"fmt"

	"github.com/lw396/WeComCopilot/pkg/sqlcipher"
	"gorm.io/gorm"
)

type SQLiteClient interface {
	OpenDB(ctx context.Context, dbName string) (tx *gorm.DB, err error)
	FindTable(ctx context.Context, tx *gorm.DB, tableName string) (sequence SQLiteSequence, err error)
}

func (s *SQLite) OpenDB(ctx context.Context, dbName string) (tx *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s/%s?_pragma_key=x'%s'", s.path, dbName, s.key)

	tx, err = gorm.Open(sqlcipher.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	return
}

type SQLiteSequence struct {
	Name string
	Seq  uint64
}

func (s *SQLite) FindTable(ctx context.Context, tx *gorm.DB, tableName string) (
	sequence SQLiteSequence, err error,
) {
	sequence = SQLiteSequence{}
	if err = tx.First(&sequence, "name = ?", tableName).Error; err != nil {
		return
	}

	return
}
