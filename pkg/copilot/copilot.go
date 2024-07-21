package copilot

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/lw396/ChatCopilot/internal/model"
	"github.com/lw396/ChatCopilot/internal/repository/gorm"
	ollama "github.com/ollama/ollama/api"
	"github.com/ollama/ollama/envconfig"
	"github.com/sashabaranov/go-openai"
)

type CopilotClient interface {
	Type() (result model.ApiType)
	Chat(ctx context.Context, msg interface{}, ch chan interface{}) (err error)
}

func NewClient(config *gorm.CopilotConfig) (CopilotClient, error) {
	switch config.ApiType {
	case model.Ollama:
		return NewOllamaClient(config), nil
	case model.Openai:
		return NewOpenaiClient(config), nil
	default:
		return nil, fmt.Errorf("unknown client type: %d", config.ApiType)
	}
}
func NewOllamaClient(config *gorm.CopilotConfig) *OllamaClient {
	ollamaHost := envconfig.Host
	if config.Url == "" {
		config.Url = net.JoinHostPort(ollamaHost.Host, ollamaHost.Port)
	}
	return &OllamaClient{
		client: ollama.NewClient(
			&url.URL{
				Scheme: ollamaHost.Scheme,
				Host:   config.Url,
			},
			http.DefaultClient,
		),
		model:       config.ModelName,
		temperature: config.Temperature,
		topP:        config.TopP,
	}
}

func NewOpenaiClient(config *gorm.CopilotConfig) *OpenaiClient {
	return &OpenaiClient{
		client:      openaiClient(config),
		model:       config.ModelName,
		temperature: config.Temperature,
		topP:        config.TopP,
	}
}

func openaiClient(config *gorm.CopilotConfig) *openai.Client {
	if config.Url == "" {
		return openai.NewClient(config.Token)
	}
	clientConfig := openai.DefaultAzureConfig(config.Token, config.Url)
	return openai.NewClientWithConfig(clientConfig)
}
