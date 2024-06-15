package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/pkg/db"
	"github.com/lw396/WeComCopilot/pkg/util"
)

const (
	SyncTaskMessageContent = "SYNC_TASK_MESSAGE_CONTENRT"
	SyncTaskUnloadedFile   = "SYNC_TASK_UNLOADED_FILE"
)

type SyncMessageTask struct {
	DBName  string
	MsgName string
	NewId   int64
	IsGroup bool
}

func (a *Service) GetCrontab() string {
	if 0 < a.task.Interval && 60 > a.task.Interval {
		return fmt.Sprintf("*/%d * * * * *", a.task.Interval)
	}
	return a.task.Crontab
}

func (a *Service) SyncMessage(ctx context.Context) (err error) {
	var params []*SyncMessageTask
	_, err = a.redis.SMembers(ctx, SyncTaskMessageContent, &params)
	if err != nil {
		return
	}

	for _, param := range params {
		go func(ctx context.Context, param *SyncMessageTask) {
			oldParam := *param
			isAlter, err := a.handleSaveMessageContent(ctx, param)
			if err != nil {
				return
			}
			if !isAlter {
				return
			}
			if err = a.redis.SUpdate(ctx, SyncTaskMessageContent, &oldParam, param); err != nil {
				return
			}
		}(ctx, param)
	}

	return
}

func (a *Service) SyncUndownloadedMessage(ctx context.Context) (err error) {
	params := []RecordUndownloadedFile{}
	found, err := a.redis.SMembers(ctx, SyncTaskUnloadedFile, &params)
	if err != nil {
		return
	}
	if found && len(params) == 0 {
		return
	}

	var deadline = time.Now().Add(-time.Minute * 10)
	for _, param := range params {
		finish, err := a.HandleUndownloadedMessage(ctx, param)
		if err != nil {
			return err
		}
		if !finish || !time.Unix(param.CreatedAt, 0).After(deadline) {
			continue
		}
		if err = a.redis.SRem(ctx, SyncTaskUnloadedFile, &param); err != nil {
			return err
		}
	}
	return
}

func (a *Service) handleSaveMessageContent(ctx context.Context, param *SyncMessageTask) (isAlter bool, err error) {
	if err = a.ConnectDB(ctx, param.DBName); err != nil {
		return
	}
	data, err := a.sqlite.GetUnsyncMessageContent(ctx, param.DBName, param.MsgName, param.NewId)
	if err != nil {
		return
	}
	if len(data) == 0 {
		return
	}

	content, err := a.HandleMessageContent(ctx, data, param.IsGroup, param.MsgName)
	if err != nil {
		return
	}
	if err = a.rep.SaveMessageContent(ctx, param.MsgName, content); err != nil {
		return
	}

	param.NewId = content[len(content)-1].LocalID
	isAlter = true
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

type InitSyncTask struct {
	UsrName string
	DBName  string
	Status  uint8
	IsGroup bool
}

func (a *Service) InitSyncTask(ctx context.Context) (err error) {
	data := []*InitSyncTask{}
	group, _, err := a.rep.GetGroupContacts(ctx, "", -1)
	if err != nil {
		return
	}
	for _, v := range group {
		data = append(data, &InitSyncTask{
			UsrName: v.UsrName,
			DBName:  v.DBName,
			Status:  v.Status,
			IsGroup: true,
		})
	}

	contact, _, err := a.rep.GetContactPersons(ctx, "", -1)
	if err != nil {
		return
	}
	for _, v := range contact {
		data = append(data, &InitSyncTask{
			UsrName: v.UsrName,
			DBName:  v.DBName,
			Status:  v.Status,
			IsGroup: false,
		})
	}

	for _, v := range data {
		msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(v.UsrName)))
		var data *mysql.MessageContent
		data, err = a.rep.GetNewMessageContent(ctx, msgName)
		if err != nil {
			if db.IsRecordNotFound(err) {
				continue
			}
			return
		}
		if err = a.ConnectMessageDB(ctx, v.DBName); err != nil {
			return
		}
		if v.Status != 1 {
			continue
		}

		param := SyncMessageTask{
			DBName:  v.DBName,
			MsgName: msgName,
			NewId:   data.LocalID,
			IsGroup: v.IsGroup,
		}
		if err = a.redis.SAdd(ctx, SyncTaskMessageContent, param); err != nil {
			return
		}
	}

	return
}

func (a *Service) AddSyncTask(ctx context.Context, msgName, dbName string, isGroup bool) (err error) {
	var data *mysql.MessageContent
	data, err = a.rep.GetNewMessageContent(ctx, msgName)
	if err != nil {
		return
	}

	param := SyncMessageTask{
		DBName:  dbName,
		MsgName: msgName,
		NewId:   data.LocalID,
		IsGroup: isGroup,
	}
	err = a.redis.SAdd(ctx, SyncTaskMessageContent, param)
	if err != nil {
		return
	}

	return
}

func (a *Service) DelSyncTask(ctx context.Context, usrName string) (err error) {
	params := []SyncMessageTask{}
	if _, err = a.redis.SMembers(ctx, SyncTaskMessageContent, &params); err != nil {
		return
	}

	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(usrName)))
	for _, param := range params {
		if param.MsgName != msgName {
			continue
		}
		if err = a.redis.SRem(ctx, SyncTaskMessageContent, param); err != nil {
			return
		}
	}
	return
}
