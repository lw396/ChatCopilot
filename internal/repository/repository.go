package repository

import (
	"context"

	db "github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
	"gorm.io/gorm"
)

type SQLiteClient interface {
	OpenDB(ctx context.Context, dbName string) (*gorm.DB, error)
	BindDB(ctx context.Context, tx *gorm.DB, dbName string)
	// Group
	GetGroupContactByNickname(ctx context.Context, nickname string) ([]*sqlite.GroupContact, error)
	// Message
	CheckMessageExistDB(ctx context.Context, tx *gorm.DB, dbName string) (*sqlite.SQLiteSequence, error)
	GetMessageContent(ctx context.Context, dbName, msgName string) ([]*sqlite.MessageContent, error)
	BindMessage(ctx context.Context, tx *gorm.DB, dbName, msgName string) error
	UnbindMessage(ctx context.Context, dbName, msgName string) error
}

type Repository interface {
	// Group
	SaveGroupContact(ctx context.Context, contact *db.GroupContact) error
	// Message
	SaveMessageContent(ctx context.Context, msgName string, content []*db.MessageContent) error
}
