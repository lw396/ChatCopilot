package service

import (
	"context"
	"strings"

	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/model"
	"github.com/lw396/WeComCopilot/internal/repository/gorm"
	ollama "github.com/ollama/ollama/api"
)

func (a *Service) AddChatCopilot(ctx context.Context, req *gorm.ChatCopilot) (err error) {
	switch req.Type {
	case model.ChatTypePerson:
		if _, err = a.rep.GetContactPersonByUsrName(ctx, req.UsrName); err != nil {
			return
		}
	case model.ChatTypeGroup:
		if _, err = a.rep.GetGroupContactByUsrName(ctx, req.UsrName); err != nil {
			return
		}
		return errors.New(errors.CodeNotSupport, "group not support")
	default:
		return errors.New(errors.CodeInvalidParam, "invalid type")
	}

	_, err = a.rep.GetPromptCuration(ctx, req.PromptID)
	if err != nil {
		return
	}

	if err = a.rep.AddChatCopilot(ctx, req); err != nil {
		return
	}
	return
}

type streamResponseFunc func(r *strings.Reader) error

func (a *Service) GetChatTips(ctx context.Context, usrname string, fn streamResponseFunc) (err error) {

	return
}
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

func (c *Service) HandleMessageFormat(message []*gorm.MessageContent) (result []ollama.Message, err error) {
	result = make([]ollama.Message, 0)
	for _, msg := range message {
		if msg.MessageType != model.MsgTypeText {
			continue
		}

		role := "user"
		if !msg.Des {
			role = "assistant"
		}
		result = append(result, ollama.Message{
			Role:    role,
			Content: msg.Content,
		})
	}

	return
}
