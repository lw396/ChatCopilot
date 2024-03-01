package sqlite

import (
	"context"
	"fmt"

	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/sqlcipher"
	"gorm.io/gorm"
)

type SQLiteClient interface {
	OpenDB(ctx context.Context, dbName string) (tx *gorm.DB, err error)
	BindDB(ctx context.Context, tx *gorm.DB, dbName string)
	// Message
	CheckMessageExistDB(ctx context.Context, tx *gorm.DB, dbName string) (sequence *repository.SQLiteSequence, err error)
	BindMessageDB(ctx context.Context, tx *gorm.DB, dbName, msgName string) (err error)
	UnbindMessage(ctx context.Context, dbName, msgName string) (err error)
	GetMessageContent(ctx context.Context, dbName, msgName string) (result []*repository.MessageContent, err error)
	// Group
	GetGroupContactByNickname(ctx context.Context, nickname string) (result []*repository.GroupContact, err error)
}

const (
	MessageDB = "Message/msg_%d.db"  // 消息库
	GroupDB   = "Group/group_new.db" // 群联系
)

type SQLite struct {
	key  string
	path string
	db   map[string]*DB
}

// BindMessageDB implements SQLiteClient.
func (s *SQLite) BindMessageDB(ctx context.Context, tx *gorm.DB, dbName string, msgName string) (err error) {
	panic("unimplemented")
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
