package crontab

import (
	"context"
)

func (s *crontabServer) InitSyncTask(ctx context.Context) (err error) {
	spec = s.service.GetCrontab()
	return s.service.InitSyncTask(ctx)
}

func (s *crontabServer) SyncMessage(ctx context.Context) (err error) {
	return s.service.SyncMessage(ctx)
}
