package copilot

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/lw396/ChatCopilot/internal/model"
	"github.com/sashabaranov/go-openai"
)

type OpenaiClient struct {
	client      *openai.Client
	model       string
	temperature float32
	topP        float32
}

func (c *OpenaiClient) Type() model.ApiType {
	return model.Openai
}

func (c *OpenaiClient) Chat(ctx context.Context, messages []Message, ch chan<- interface{}) error {
	if len(messages) == 0 {
		close(ch)
		return nil
	}

	reqMessages := make([]openai.ChatCompletionMessage, 0, len(messages))
	for _, message := range messages {
		reqMessages = append(reqMessages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	stream, err := c.client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:       c.model,
		Messages:    reqMessages,
		Temperature: c.temperature,
		TopP:        c.topP,
		Stream:      true,
	})
	if err != nil {
		return err
	}
	defer stream.Close()
	defer close(ch)

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("stream error: %w", err)
		}

		ch <- response
	}
}
