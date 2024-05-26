package service

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/lw396/WeComCopilot/internal/errors"
	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
	"github.com/lw396/WeComCopilot/pkg/db"
	"github.com/lw396/WeComCopilot/pkg/util"
	"gorm.io/gorm"
)

type MessageInfo struct {
	UserName string `json:"user_name"`
	Seq      uint64 `json:"seq"`
	DBName   string `json:"db_name"`
}

func (a *Service) ScanMessage(ctx context.Context, userName string) (result *MessageInfo, err error) {
	var dbName string
	var seq *sqlite.SQLiteSequence
	var name string = "Chat_" + hex.EncodeToString(util.Md5([]byte(userName)))
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
		DBName:   dbName,
		UserName: userName,
		Seq:      seq.Seq,
	}
	return
}

func (a *Service) GetMessageContent(ctx context.Context, usrName string, offset int) (result []*mysql.MessageContent, err error) {
	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(usrName)))
	result, err = a.rep.GetMessageContentList(ctx, msgName, offset)
	if err != nil {
		return
	}

	return
}

func (a *Service) convertMessageContent(msg []*sqlite.MessageContent) (result []*mysql.MessageContent) {
	result = make([]*mysql.MessageContent, 0)
	for _, v := range msg {
		result = append(result, &mysql.MessageContent{
			LocalID:     v.MesLocalID,
			SvrID:       v.MesSvrID,
			CreateTime:  v.MsgCreateTime,
			Content:     v.MsgContent,
			Status:      v.MsgStatus,
			ImgStatus:   v.MsgImgStatus,
			MessageType: v.MessageType,
			Des:         v.MesDes,
			Source:      v.MsgSource,
			VoiceText:   v.MsgVoiceText,
			Seq:         v.MsgSeq,
		})
	}
	return
}
