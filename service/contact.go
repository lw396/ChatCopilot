package service

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
)

type ContactPerson struct {
	UsrName    string `json:"usr_name"`
	Nickname   string `json:"nickname"`
	Remark     string `json:"remark"`
	HeadImgUrl string `json:"head_img_url"`
	DBName     string `json:"db_name,omitempty"`
	Status     uint8  `json:"status,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
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
			HeadImgUrl: row.HeadImgUrl,
		})
	}
	return
}
