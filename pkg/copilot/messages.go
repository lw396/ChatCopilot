package copilot

import (
	"context"
	"fmt"

	ollama "github.com/ollama/ollama/api"
)

func (c *CopilotClient) Chat(ctx context.Context, messages []ollama.Message, ch chan interface{}) (err error) {
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

	errCh := make(chan error)
	respFunc := func(resp ollama.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		ch <- resp
		return nil
	}
	go func() {
		errCh <- nil
		if err = c.ollama.Chat(ctx, req, respFunc); err != nil {
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
