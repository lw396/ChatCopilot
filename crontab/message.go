package crontab

import (
	"context"
)

func (s *crontabServer) SyncMessage(ctx context.Context) (err error) {
	err = s.service.SyncMessage(ctx)
	if err != nil {
		return
	}
	return
}
