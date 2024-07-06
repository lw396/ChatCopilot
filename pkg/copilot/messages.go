package copilot

import (
	"context"

	ollama "github.com/ollama/ollama/api"
)

func (c *CopilotClient) Chat(ctx context.Context, messages []ollama.Message) (err error) {
	stream := true
	req := &ollama.ChatRequest{
		Model:    c.model,
		Messages: messages,
		Stream:   &stream,
		Options: map[string]interface{}{
			"temperature": c.temperature,
			"top_p":       c.topP,
		},
	}

	chat := ollama.Message{}
	respFunc := func(resp ollama.ChatResponse) error {
		chat = ollama.Message{
			Role:    resp.Message.Role,
			Content: chat.Content + resp.Message.Content,
		}

		return nil
	}

	if err = c.ollama.Chat(ctx, req, respFunc); err != nil {
		return
	}

	return
}
