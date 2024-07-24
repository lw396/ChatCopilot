package service

import (
	"context"
	"encoding/hex"

	"github.com/lw396/ChatCopilot/internal/errors"
	"github.com/lw396/ChatCopilot/internal/model"
	"github.com/lw396/ChatCopilot/internal/repository/gorm"
	"github.com/lw396/ChatCopilot/pkg/util"
	ollama "github.com/ollama/ollama/api"
	"github.com/sashabaranov/go-openai"
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
	msgName := "Chat_" + hex.EncodeToString(util.Md5([]byte(usrname)))
	messages, err := a.rep.GetMessageContentList(ctx, msgName, -1, -1)
	if err != nil {
		return
	}

	message, err := a.HandleMessageFormat(ctx, messages, copilot)
	if err != nil {
		return
	}

	err = a.copilot.Chat(ctx, message, ch)
	if err != nil {
		return
	}

	return
}

func (a *Service) HandleMessageFormat(ctx context.Context, messages []*gorm.MessageContent, copilot *gorm.ChatCopilot) (
	result interface{}, err error) {
	switch a.copilot.Type() {
	case model.Ollama:
		result = a.HandleOllamaMessage(messages, copilot)
	case model.Openai:
		result = a.HandleOpanaiMessage(messages, copilot)
	default:
		err = errors.New(errors.CodeNotSupport, "not support")
	}
	return
}

func (a *Service) HandleOllamaMessage(messages []*gorm.MessageContent, copilot *gorm.ChatCopilot) (
	result []ollama.Message) {
	result = append(result, ollama.Message{
		Role:    "system",
		Content: copilot.Prompt.Prompt,
	})

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

func (a *Service) HandleOpanaiMessage(messages []*gorm.MessageContent, copilot *gorm.ChatCopilot) (
	result []openai.ChatCompletionMessage) {
	result = append(result, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: copilot.Prompt.Prompt,
	})

	for _, msg := range messages {
		if msg.MessageType != model.MsgTypeText {
			continue
		}

		role := openai.ChatMessageRoleUser
		if !msg.Des {
			role = openai.ChatMessageRoleAssistant
		}
		result = append(result, openai.ChatCompletionMessage{
			Role:    role,
			Content: msg.Content,
		})
	}
	return
}
