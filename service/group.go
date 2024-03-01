package service

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
)

type GroupContact struct {
	UsrName         string `json:"usr_name"`
	Nickname        string `json:"nickname"`
	HeadImgUrl      string `json:"head_img_url"`
	ChatRoomMemList string `json:"member_list"`
}

func (a *Service) GetGroupContact(ctx context.Context, nickname string) (result []*GroupContact, err error) {
	if err = a.ConnectGroup(ctx); err != nil {
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

func (a *Service) ConnectGroup(ctx context.Context) (err error) {
	tx, err := a.sqlite.OpenDB(ctx, sqlite.GroupDB)
	if err != nil {
		return
	}

	if tx != nil {
		a.sqlite.BindDB(ctx, tx, sqlite.GroupDB)
	}
	return
}
