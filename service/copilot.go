package service

import (
	"context"
	"encoding/hex"

	"github.com/lw396/WeComCopilot/internal/errors"
	"github.com/lw396/WeComCopilot/internal/model"
	"github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/pkg/util"
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

func (a *Service) GetChatTips(ctx context.Context, usrname string, ch chan interface{}) (err error) {
	copilot, err := a.rep.GetChatCopilotByUsrName(ctx, usrname)
	if err != nil {
		return
	}
	message, err := a.HandleMessageFormat(ctx, usrname)
	if err != nil {
		return
	}
	promptMessages := []ollama.Message{{
		Role:    "system",
		Content: copilot.Prompt.Prompt,
	}}
	promptMessages = append(promptMessages, message...)
	err = a.copilot.Chat(ctx, promptMessages, ch)
	if err != nil {
		return
	}

	return
}

func (a *Service) HandleMessageFormat(ctx context.Context, usrname string) (result []ollama.Message, err error) {
	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(usrname)))
	messages, err := a.rep.GetMessageContentList(ctx, msgName, -1, -1)
	if err != nil {
		return
	}

	result = make([]ollama.Message, 0)
	for _, msg := range messages {
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
