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

type ContactPerson struct {
	Id         uint64    `json:"id"`
	UsrName    string    `json:"usr_name"`
	Nickname   string    `json:"nickname"`
	Remark     string    `json:"remark"`
	HeadImgUrl string    `json:"head_img_url"`
	Sex        int64     `json:"sex"`
	Type       int64     `json:"type"`
	DBName     string    `json:"db_name,omitempty"`
	Status     uint8     `json:"status,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (a *Service) GetContactPersonByNickname(ctx context.Context, nickname string) (result []*ContactPerson, err error) {
	if err = a.ConnectDB(ctx, sqlite.ContactDB); err != nil {
		return
	}
	contact, err := a.sqlite.GetContactPersonByNickname(ctx, nickname)
	if err != nil {
		return
	}
	for _, row := range contact {
		result = append(result, &ContactPerson{
			UsrName:    row.UsrName,
			Nickname:   row.Nickname,
			Remark:     row.Remark,
			Sex:        row.Sex,
			Type:       row.Type,
			HeadImgUrl: row.HeadImgUrl,
		})
	}
	return
}

func (a *Service) GetContactPersonByUsrname(ctx context.Context, usrname string) (result *ContactPerson, err error) {
	if err = a.ConnectDB(ctx, sqlite.ContactDB); err != nil {
		return
	}

	var contact *sqlite.ContactPerson
	contact, err = a.sqlite.GetContactPersonByUsrname(ctx, usrname)
	if err != nil {
		return
	}

	result = &ContactPerson{
		UsrName:    contact.UsrName,
		Nickname:   contact.Nickname,
		Remark:     contact.Remark,
		HeadImgUrl: contact.HeadImgUrl,
	}
	return
}

func (a *Service) GetContactPersonList(ctx context.Context, offset int, nickname string) (result []*ContactPerson, totle int64, err error) {
	contact, totle, err := a.rep.GetContactPersons(ctx, nickname, offset)
	if err != nil {
		return
	}
	for _, v := range contact {
		result = append(result, &ContactPerson{
			Id:         v.ID,
			UsrName:    v.UsrName,
			Nickname:   v.Nickname,
			HeadImgUrl: v.HeadImgUrl,
			Sex:        v.Sex,
			Type:       v.Type,
			Remark:     v.Remark,
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
		})
	}
	return
}

func (a *Service) SaveContactPerson(ctx context.Context, data *ContactPerson) (err error) {
	if err = a.ConnectMessageDB(ctx, data.DBName); err != nil {
		return
	}
	if _, err = a.rep.GetContactPersonByUsrName(ctx, data.UsrName); err != gorm.ErrRecordNotFound {
		if err == nil {
			err = errors.New(errors.CodeAuthMessageFound, "contact person already exist")
		}
		return
	}

	if err = a.rep.SaveContactPerson(ctx, &mysql.ContactPerson{
		UsrName:    data.UsrName,
		Nickname:   data.Nickname,
		Remark:     data.Remark,
		HeadImgUrl: data.HeadImgUrl,
		Type:       data.Type,
		DBName:     data.DBName,
		Status:     1,
	}); err != nil {
		return
	}

	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(data.UsrName)))
	messages, err := a.sqlite.GetMessageContent(ctx, data.DBName, msgName)
	if err != nil {
		return
	}

	if err = a.rep.CreateMessageContentTable(ctx, msgName); err != nil {
		return
	}

	content, err := a.convertMessageContent(ctx, messages, false)
	if err != nil {
		return
	}
	if err = a.rep.SaveMessageContent(ctx, msgName, content); err != nil {
		return
	}

	if err := a.AddSyncTask(ctx, msgName, data.DBName, false); err != nil {
		return err
	}

	return
}

func (a *Service) DelContactPerson(ctx context.Context, usrName string) (err error) {
	if err = a.rep.DelContactPersonByUsrName(ctx, usrName); err != nil {
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
