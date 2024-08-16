package copilot

import (
	"context"
	"fmt"

	"github.com/lw396/ChatCopilot/internal/model"
	ollama "github.com/ollama/ollama/api"
)

type OllamaClient struct {
	client      *ollama.Client
	model       string
	temperature float32
	topP        float32
}

func (c *OllamaClient) Type() (result model.ApiType) {
	return model.Ollama
}

func (c *OllamaClient) Chat(ctx context.Context, msg interface{}, ch chan interface{}) (err error) {
	message := msg.([]ollama.Message)
	stream := true

	req := &ollama.ChatRequest{
		Model:    c.model,
		Messages: message,
		Stream:   &stream,
		Options: map[string]interface{}{
			"temperature": c.temperature,
			"top_p":       c.topP,
		},
	}

	errCh := make(chan error)
	respFunc := func(resp ollama.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		ch <- resp
		return nil
	}
	go func() {
		errCh <- nil
		if err = c.client.Chat(ctx, req, respFunc); err != nil {
			errCh <- err
			return
		}
		defer close(ch)
	}()
	if err := <-errCh; err != nil {
		return err
	}

	return
}
