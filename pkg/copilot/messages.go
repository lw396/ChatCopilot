package copilot

import (
	"context"
	"fmt"

	ollama "github.com/ollama/ollama/api"
)

type CopilotClient struct {
	ollama      *ollama.Client
	model       string
	temperature float32
	topP        float32
	stream      bool
}

func (c *CopilotClient) Chat(ctx context.Context, messages []ollama.Message) (err error) {
	req := &ollama.ChatRequest{
		Model:    c.model,
		Messages: messages,
		Stream:   &c.stream,
		Options: map[string]interface{}{
			"temperature": c.temperature,
			"top_p":       c.topP,
		},
	}

	chat := ollama.Message{}
	channel := make(chan ollama.Message)
	respFunc := func(resp ollama.ChatResponse) error {
		chat = ollama.Message{
			Role:    resp.Message.Role,
			Content: chat.Content + resp.Message.Content,
		}
		channel <- chat

		return nil
	}

	if err = c.ollama.Chat(ctx, req, respFunc); err != nil {
		return
	}

	fmt.Println(<-channel)
	return
}
