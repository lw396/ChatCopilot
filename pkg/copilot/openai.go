package copilot

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenaiClient struct {
	client      *openai.Client
	model       string
	temperature float32
	topP        float32
}

func (c *OpenaiClient) Chat(ctx context.Context, msg interface{}, ch chan interface{}) (err error) {
	c.client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Messages: []openai.ChatCompletionMessage{},
	})
	return
}
