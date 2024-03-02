package sqlite

import (
	"context"
	"fmt"

	"github.com/lw396/WeComCopilot/pkg/sqlcipher"
	"gorm.io/gorm"
)

const (
	MessageDB = "Message/msg_%d.db"  // 消息库
	GroupDB   = "Group/group_new.db" // 群联系
)

type SQLite struct {
	key  string
	path string
	db   map[string]*DB
}

type DB struct {
	tx      *gorm.DB
	msgName []string
}

func NewSQLiteClient(key, path string) *SQLite {
	return &SQLite{
		key:  key,
		path: path,
		db:   make(map[string]*DB),
	}
}

func (s *SQLite) OpenDB(ctx context.Context, dbName string) (tx *gorm.DB, err error) {
	if s.db[dbName] != nil {
		return
	}
	dsn := fmt.Sprintf("%s/%s?_pragma_key=x'%s'", s.path, dbName, s.key)
	return gorm.Open(sqlcipher.Open(dsn), &gorm.Config{})
}

func (s *SQLite) BindDB(ctx context.Context, tx *gorm.DB, dbName string) {
	s.db[dbName] = &DB{tx: tx, msgName: []string{}}
}
