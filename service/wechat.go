package service

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/internal/repository/sqlite"
)

func (a *Service) GetGroupContact(ctx context.Context, nickname string) (result []*repository.GroupContact, err error) {
	if err = a.ConnectGroup(ctx); err != nil {
		return
	}

	result, err = a.sqlite.GetGroupContactByNickname(ctx, nickname)
	if err != nil {
		return
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
