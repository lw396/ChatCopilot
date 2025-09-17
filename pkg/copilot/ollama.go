package copilot

import (
	"context"

	"github.com/lw396/ChatCopilot/internal/model"
	ollama "github.com/ollama/ollama/api"
)

type OllamaClient struct {
	client      *ollama.Client
	model       string
	temperature float32
	topP        float32
}

func (c *OllamaClient) Type() model.ApiType {
	return model.Ollama
}

func (c *OllamaClient) Chat(ctx context.Context, messages []Message, ch chan<- interface{}) error {
	if len(messages) == 0 {
		close(ch)
		return nil
	}

	ollamaMessages := make([]ollama.Message, 0, len(messages))
	for _, message := range messages {
		ollamaMessages = append(ollamaMessages, ollama.Message{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	stream := true
	req := &ollama.ChatRequest{
		Model:    c.model,
		Messages: ollamaMessages,
		Stream:   &stream,
		Options: map[string]interface{}{
			"temperature": c.temperature,
			"top_p":       c.topP,
		},
	}

	defer close(ch)

	return c.client.Chat(ctx, req, func(resp ollama.ChatResponse) error {
		ch <- resp
		return nil
	})
}
