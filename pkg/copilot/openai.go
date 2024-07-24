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

func (c *OpenaiClient) Type() (result model.ApiType) {
	return model.Openai
}

func (c *OpenaiClient) Chat(ctx context.Context, msg interface{}, ch chan interface{}) (err error) {
	// messages := msg.([]openai.ChatCompletionMessage)
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "Hello!",
		},
	}
	fmt.Println(messages)
	stream, err := c.client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:       c.model,
		Messages:    messages,
		Temperature: c.temperature,
		TopP:        c.topP,
		Stream:      true,
	})
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		var response openai.ChatCompletionStreamResponse
		response, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			err = errors.New("Stream error: " + err.Error())
			return
		}

		fmt.Println(response.Choices[0].Delta.Content)
		ch <- response
	}
}
