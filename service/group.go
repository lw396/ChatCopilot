package service

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/lw396/WeComCopilot/internal/errors"
	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
	"github.com/lw396/WeComCopilot/pkg/util"
	"gorm.io/gorm"
)

type GroupContact struct {
	Id              uint64    `json:"id"`
	UsrName         string    `json:"usr_name"`
	Nickname        string    `json:"nickname"`
	HeadImgUrl      string    `json:"head_img_url"`
	ChatRoomMemList string    `json:"member_list"`
	DBName          string    `json:"db_name,omitempty"`
	Status          uint8     `json:"status,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

func (a *Service) GetGroupContactByNickname(ctx context.Context, nickname string) (result []*GroupContact, err error) {
	if err = a.ConnectDB(ctx, sqlite.GroupDB); err != nil {
		return
	}
	contact, err := a.sqlite.GetGroupContactByNickname(ctx, nickname)
	if err != nil {
		return
	}
	for _, row := range contact {
		result = append(result, &GroupContact{
			UsrName:         row.UsrName,
			Nickname:        row.Nickname,
			HeadImgUrl:      row.HeadImgUrl,
			ChatRoomMemList: row.ChatRoomMemList,
		})
	}
	return
}

func (a *Service) GetGroupContactByUsrname(ctx context.Context, usrname string) (result *GroupContact, err error) {
	if err = a.ConnectDB(ctx, sqlite.GroupDB); err != nil {
		return
	}

	var group *sqlite.GroupContact
	group, err = a.sqlite.GetGroupContactByUsrname(ctx, usrname)
	if err != nil {
		return
	}

	result = &GroupContact{
		UsrName:         group.UsrName,
		Nickname:        group.Nickname,
		HeadImgUrl:      group.HeadImgUrl,
		ChatRoomMemList: group.ChatRoomMemList,
	}
	return
}

func (a *Service) GetGroupContactList(ctx context.Context, offset int, nickname string) (result []*GroupContact, totle int64, err error) {
	group, totle, err := a.rep.GetGroupContacts(ctx, nickname, offset)
	if err != nil {
		return
	}
	for _, v := range group {
		result = append(result, &GroupContact{
			Id:              v.ID,
			UsrName:         v.UsrName,
			Nickname:        v.Nickname,
			HeadImgUrl:      v.HeadImgUrl,
			ChatRoomMemList: v.ChatRoomMemList,
			Status:          v.Status,
			CreatedAt:       v.CreatedAt,
		})
	}
	return
}

func (a *Service) SaveGroupContact(ctx context.Context, data *GroupContact) (err error) {
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
		Status:          1,
	}); err != nil {
		return
	}

	if err = a.rep.CreateMessageContentTable(ctx, msgName); err != nil {
		return
	}

	content, err := a.HandleMessageContent(ctx, messages, true, msgName)
	if err != nil {
		return
	}
	if err = a.rep.SaveMessageContent(ctx, msgName, content); err != nil {
		return
	}

	if err = a.AddSyncTask(ctx, msgName, data.DBName, true); err != nil {
		return
	}

	return
}

func (a *Service) DelGroupContact(ctx context.Context, usrName string) (err error) {
	if err = a.rep.DelGroupContactByUsrName(ctx, usrName); err != nil {
		return
	}

	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(usrName)))
	err = a.rep.DelMessageContentTable(ctx, msgName)
	if err != nil {
		return
	}

	if err = a.DelSyncTask(ctx, usrName); err != nil {
		return
	}

	return
}

func (a *Service) ConnectDB(ctx context.Context, dbName string) (err error) {
	tx, err := a.sqlite.OpenDB(ctx, dbName)
	if err != nil {
		return
	}
	if tx != nil {
		a.sqlite.BindDB(ctx, tx, dbName)
	}
	return
}
