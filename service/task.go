package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/pkg/db"
	"github.com/lw396/WeComCopilot/pkg/util"
)

const (
	SyncTaskMessageContent = "SYNC_TASK_MESSAGE_CONTENRT"
	SyncTaskUnloadedFile   = "SYNC_TASK_UNLOADED_FILE"
)

type SyncMessageTaskParam struct {
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
	var params []SyncMessageTaskParam
	_, err = a.redis.Get(ctx, SyncTaskMessageContent, &params)
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}
	newParams := make([]SyncMessageTaskParam, len(params))
	for i, param := range params {
		wg.Add(1)
		go func(ctx context.Context, param SyncMessageTaskParam, i int, newParams []SyncMessageTaskParam) {
			defer wg.Done()
			newParams[i], err = a.handleSaveMessageContent(ctx, param)
			if err != nil {
				return
			}
		}(ctx, param, i, newParams)
	}
	wg.Wait()

	if err = a.redis.Set(ctx, SyncTaskMessageContent, newParams, 0); err != nil {
		return
	}
	return
}

func (a *Service) SyncUndownloadedMessage(ctx context.Context) (err error) {
	params := []RecordUndownloadedFileParam{}
	found, err := a.redis.Get(ctx, SyncTaskUnloadedFile, &params)
	if err != nil {
		return
	}
	if found && len(params) == 0 {
		return
	}

	var now = time.Now()
	for _, param := range params {
		if !param.CreatedAt.After(now) {
			continue
		}
		finish, err := a.HandleUndownloadedMessage(ctx, param)
		if err != nil {
			return err
		}
		if !finish {
			params = append(params, param)
		}
	}

	if err = a.redis.Set(ctx, SyncTaskUnloadedFile, params, 10*time.Minute); err != nil {
		return
	}
	return
}

func (a *Service) handleSaveMessageContent(ctx context.Context, param SyncMessageTaskParam) (newParam SyncMessageTaskParam, err error) {
	newParam = param
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

	newParam.NewId = content[len(content)-1].LocalID
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

type InitSyncTaskParam struct {
	UsrName string
	DBName  string
	Status  uint8
	IsGroup bool
}

func (a *Service) InitSyncTask(ctx context.Context) (err error) {
	data := []*InitSyncTaskParam{}
	group, _, err := a.rep.GetGroupContacts(ctx, "", -1)
	if err != nil {
		return
	}
	for _, v := range group {
		data = append(data, &InitSyncTaskParam{
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
		data = append(data, &InitSyncTaskParam{
			UsrName: v.UsrName,
			DBName:  v.DBName,
			Status:  v.Status,
			IsGroup: false,
		})
	}

	param := []SyncMessageTaskParam{}
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
		param = append(param, SyncMessageTaskParam{
			DBName:  v.DBName,
			MsgName: msgName,
			NewId:   data.LocalID,
			IsGroup: v.IsGroup,
		})
	}

	err = a.redis.Set(ctx, SyncTaskMessageContent, param, 0)
	if err != nil {
		return
	}

	return
}

func (a *Service) AddSyncTask(ctx context.Context, msgName, dbName string, isGroup bool) (err error) {
	var data *mysql.MessageContent
	data, err = a.rep.GetNewMessageContent(ctx, msgName)
	if err != nil {
		return
	}
	if err = a.ConnectMessageDB(ctx, dbName); err != nil {
		return
	}

	param := []SyncMessageTaskParam{}
	_, err = a.redis.Get(ctx, SyncTaskMessageContent, &param)
	if err != nil {
		return
	}

	param = append(param, SyncMessageTaskParam{
		DBName:  dbName,
		MsgName: msgName,
		NewId:   data.LocalID,
		IsGroup: isGroup,
	})
	err = a.redis.Set(ctx, SyncTaskMessageContent, param, 0)
	if err != nil {
		return
	}

	return
}

func (a *Service) DelSyncTask(ctx context.Context, usrName string) (err error) {
	_param := []SyncMessageTaskParam{}
	_, err = a.redis.Get(ctx, SyncTaskMessageContent, &_param)
	if err != nil {
		return
	}

	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(usrName)))
	param := []SyncMessageTaskParam{}
	for _, p := range _param {
		if p.MsgName == msgName {
			continue
		}
		param = append(param, p)
	}
	err = a.redis.Set(ctx, SyncTaskMessageContent, param, 0)
	if err != nil {
		return
	}
	return
}
