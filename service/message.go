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

func (a *Service) SaveMessageContent(ctx context.Context, data *GroupContact) (err error) {
	if err = a.ConnectMessageDB(ctx, data.DBName); err != nil {
		return
	}
	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(data.UsrName)))
	messages, err := a.sqlite.GetMessageContent(ctx, data.DBName, msgName)
	if err != nil {
		return
	}

	if _, err = a.rep.GetGroupContactByUsrName(ctx, data.UsrName); err != gorm.ErrRecordNotFound {
		if err == nil {
			err = errors.New(errors.CodeAuthMessageFound, "group already exist")
		}
		return
	}
	if err = a.rep.SaveGroupContact(ctx, &mysql.GroupContact{
		UsrName:         data.UsrName,
		Nickname:        data.Nickname,
		HeadImgUrl:      data.HeadImgUrl,
		ChatRoomMemList: data.ChatRoomMemList,
		DBName:          data.DBName,
	}); err != nil {
		return
	}

	if err = a.rep.CreateMessageContentTable(ctx, msgName); err != nil {
		return
	}

	content := a.convertMessageContent(messages)
	if err = a.rep.SaveMessageContent(ctx, msgName, content); err != nil {
		return
	}

	// Update sync task
	go func(ctx context.Context) {
		if err := a.InitSyncTask(ctx); err != nil {
			a.logger.Errorf("update sync task failed, err: %v", err)
		}
	}(ctx)

	return
}

type SyncMessageTaskParam struct {
	DBName  string
	MsgName string
	NewId   int64
}

const (
	SyncTaskCacheKey = "SYNC_TASK_CACHE_PARAM"
)

func (a *Service) InitSyncTask(ctx context.Context) (err error) {
	group, err := a.rep.GetGroupContacts(ctx)
	if err != nil {
		return
	}

	param := make([]SyncMessageTaskParam, 0)
	for _, v := range group {
		msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(v.UsrName)))
		var data *mysql.MessageContent
		data, err = a.rep.GetNewMessageContent(ctx, msgName)
		if err != nil {
			return
		}
		if err = a.ConnectMessageDB(ctx, v.DBName); err != nil {
			return
		}
		param = append(param, SyncMessageTaskParam{
			DBName:  v.DBName,
			MsgName: msgName,
			NewId:   data.LocalID,
		})
	}
	err = a.redis.Set(ctx, SyncTaskCacheKey, param, 0)
	if err != nil {
		return
	}

	return
}

func (a *Service) SyncMessage(ctx context.Context) (err error) {
	var params []SyncMessageTaskParam
	_, err = a.redis.Get(ctx, SyncTaskCacheKey, &params)
	if err != nil {
		return
	}

	for i, param := range params {
		go func(ctx context.Context, param SyncMessageTaskParam, i int) {
			data, err := a.sqlite.GetUnsyncMessageContent(ctx, param.DBName, param.MsgName, param.NewId)
			if err != nil {
				if db.IsRecordNotFound(err) {
					err = nil
					return
				}
				return
			}
			content := a.convertMessageContent(data)
			if err = a.rep.SaveMessageContent(ctx, param.MsgName, content); err != nil {
				return
			}
			params[i].NewId = content[len(content)-1].LocalID
		}(ctx, param, i)
	}
	return
}

func (a *Service) ConnectMessageDB(ctx context.Context, dbName string) (err error) {
	tx, err := a.sqlite.OpenDB(ctx, dbName)
	if err != nil {
		return
	}

	if err = a.sqlite.BindMessageDB(ctx, tx, dbName); err != nil {
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
