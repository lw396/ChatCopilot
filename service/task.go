package service

import (
	"context"
	"encoding/hex"

	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/pkg/db"
	"github.com/lw396/WeComCopilot/pkg/util"
)

const (
	SyncTaskCacheKey = "SYNC_TASK_CACHE_PARAM"
)

type SyncMessageTaskParam struct {
	DBName  string
	MsgName string
	NewId   int64
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

func (a *Service) InitSyncTask(ctx context.Context) (err error) {
	group, _, err := a.rep.GetGroupContacts(ctx, 0)
	if err != nil {
		return
	}

	param := make([]SyncMessageTaskParam, 0)
	for _, v := range group {
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
		})
	}

	err = a.redis.Set(ctx, SyncTaskCacheKey, param, 0)
	if err != nil {
		return
	}

	return
}

func (a *Service) AddSyncTask(ctx context.Context, msgName, dbName string) (err error) {
	var data *mysql.MessageContent
	data, err = a.rep.GetNewMessageContent(ctx, msgName)
	if err != nil {
		return
	}
	if err = a.ConnectMessageDB(ctx, dbName); err != nil {
		return
	}

	param := make([]SyncMessageTaskParam, 0)
	_, err = a.redis.Get(ctx, SyncTaskCacheKey, &param)
	if err != nil {
		return
	}

	param = append(param, SyncMessageTaskParam{
		DBName:  dbName,
		MsgName: msgName,
		NewId:   data.LocalID,
	})
	err = a.redis.Set(ctx, SyncTaskCacheKey, param, 0)
	if err != nil {
		return
	}

	return
}
