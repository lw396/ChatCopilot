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
	GetGroupContactByUsrname(ctx context.Context, usrname string) (*sqlite.GroupContact, error)

	// Contact
	GetContactPersonByNickname(ctx context.Context, nickname string) ([]*sqlite.ContactPerson, error)
	GetContactPersonByUsrname(ctx context.Context, usrname string) (*sqlite.ContactPerson, error)

	// Message
	BindMessageDB(ctx context.Context, tx *gorm.DB, dbName string) error
	UnbindMessageDB(ctx context.Context, dbName string)
	CheckMessageExistDB(ctx context.Context, tx *gorm.DB, dbName string) (*sqlite.SQLiteSequence, error)
	GetMessageContent(ctx context.Context, dbName, msgName string) ([]*sqlite.MessageContent, error)
	GetUnsyncMessageContent(ctx context.Context, dbName, msgName string, newId int64) ([]*sqlite.MessageContent, error)

	// Hink
	GetHinkMediaByMediaMd5(ctx context.Context, mediaMd5 string) (*sqlite.HlinkMediaRecord, error)
}

type Repository interface {
	// Group
	SaveGroupContact(ctx context.Context, contact *db.GroupContact) error
	GetGroupContacts(ctx context.Context, nickname string, offset int) ([]*db.GroupContact, int64, error)
	DelGroupContactByUsrName(ctx context.Context, usrName string) error
	GetGroupContactByUsrName(ctx context.Context, usrName string) (*db.GroupContact, error)

	// Contact
	SaveContactPerson(ctx context.Context, contact *db.ContactPerson) error
	GetContactPersons(ctx context.Context, nickname string, offset int) ([]*db.ContactPerson, int64, error)
	DelContactPersonByUsrName(ctx context.Context, usrname string) error
	GetContactPersonByUsrName(ctx context.Context, usrName string) (*db.ContactPerson, error)

	// Message
	CreateMessageContentTable(ctx context.Context, msgName string) error
	SaveMessageContent(ctx context.Context, msgName string, content []*db.MessageContent) error
	GetNewMessageContent(ctx context.Context, msgName string) (*db.MessageContent, error)
	DelMessageContentTable(ctx context.Context, msgName string) error
	GetMessageContentList(ctx context.Context, msgName string, offset int) ([]*db.MessageContent, error)
}
