package copilot

import (
	json "github.com/json-iterator/go"

	"github.com/lw396/WeComCopilot/internal/model"
	db "github.com/lw396/WeComCopilot/internal/repository/gorm"
	ollama "github.com/ollama/ollama/api"
)

type Copilot struct {
	Client   *ollama.Client
	Messages []byte
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (c *Copilot) HandleMessageToPrompt(message []*db.MessageContent) (result []byte, err error) {
	messages := make([]Messages, 0)
	for _, msg := range message {
		if msg.MessageType != model.MsgTypeText {
			continue
		}

		role := "user"
		if !msg.Des {
			role = "assistant"
		}
		messages = append(messages, Messages{
			Role:    role,
			Content: msg.Content,
		})
	}

	result, err = json.Marshal(messages)
	if err != nil {
		return
	}
	return
}
