package service

import (
	"context"

	"github.com/lw396/WeComCopilot/internal/model"
	mysql "github.com/lw396/WeComCopilot/internal/repository/gorm"

	ollama "github.com/ollama/ollama/api"
)

func (a *Service) InitCopilot(ctx context.Context, msgName string) (err error) {
	messages, err := a.rep.GetMessageContentList(ctx, msgName, 1, -1)
	if err != nil {
		return
	}

	data, err := a.HandleMessageFormat(messages)
	if err != nil {
		return
	}

	_ = data

	return
}

func (c *Service) HandleMessageFormat(message []*mysql.MessageContent) (result []ollama.Message, err error) {
	messages := make([]ollama.Message, 0)
	for _, msg := range message {
		if msg.MessageType != model.MsgTypeText {
			continue
		}

		role := "user"
		if !msg.Des {
			role = "assistant"
		}
		messages = append(messages, ollama.Message{
			Role:    role,
			Content: msg.Content,
		})
	}

	return
}
