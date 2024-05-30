package sqlite

import (
	"context"
	"fmt"

	"github.com/lw396/WeComCopilot/pkg/sqlcipher"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	MessageDB = "Message/msg_%d.db"         // 消息库
	GroupDB   = "Group/group_new.db"        // 群联系
	ContactDB = "Contact/wccontact_new2.db" // 联系人
	HlinkDB   = "Hardlink/hldata.db"        // 文件路径

	FavArchive = "Stickers/fav.archive" // 收藏表情包
)

type SQLite struct {
	key  string
	path string
	db   map[string]*gorm.DB
}

func NewSQLite(key, path string) *SQLite {
	return &SQLite{
		key:  key,
		path: path,
		db:   make(map[string]*gorm.DB),
	}
}

func (s *SQLite) OpenDB(ctx context.Context, dbName string) (tx *gorm.DB, err error) {
	if s.db[dbName] != nil {
		return s.db[dbName], err
	}
	dsn := fmt.Sprintf("%s/%s?_pragma_key=x'%s'", s.path, dbName, s.key)
	return gorm.Open(sqlcipher.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func (s *SQLite) BindDB(ctx context.Context, tx *gorm.DB, dbName string) {
	if s.db[dbName] == nil {
		s.db[dbName] = tx
	}
}
