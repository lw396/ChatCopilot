package service

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
	"github.com/lw396/WeComCopilot/pkg/db"
	"github.com/lw396/WeComCopilot/pkg/util"
	"gorm.io/gorm"
)

type MessageInfo struct {
	Name   string `json:"name"`
	Seq    uint64 `json:"seq"`
	DBName string `json:"db_name"`
}

func (a *Service) ScanMessage(ctx context.Context, userName string) (result *MessageInfo, err error) {
	var dbName string
	var seq *repository.SQLiteSequence
	name := "Chat_" + hex.EncodeToString(util.Md5([]byte(userName)))
	fmt.Println(name)
	for i := 0; i < 10; i++ {
		dbName = fmt.Sprintf(sqlite.MessageDB, i)
		var tx *gorm.DB
		if tx, err = a.sqlite.OpenDB(ctx, dbName); err != nil {
			return
		}
		if seq, err = a.sqlite.CheckMessageExistDB(ctx, tx, name); err != nil {
			if !db.IsRecordNotFound(err) {
				return
			}
			continue
		}
		break
	}
	if seq == nil {
		err = errors.New(errors.CodeDB, "not found message")
		return
	}
	result = &MessageInfo{
		DBName: dbName,
		Name:   seq.Name,
		Seq:    seq.Seq,
	}
	return
}
