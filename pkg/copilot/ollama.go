package copilot

import (
	"net"
	"net/http"
	"net/url"

	"github.com/lw396/ChatCopilot/internal/repository/gorm"
	ollama "github.com/ollama/ollama/api"
	"github.com/ollama/ollama/envconfig"
)

type CopilotClient struct {
	ollama      *ollama.Client
	model       string
	temperature float32
	topP        float32
}

func NewClient(config *gorm.CopilotConfig) *CopilotClient {
	ollamaHost := envconfig.Host
	if config.Url == "" {
		config.Url = net.JoinHostPort(ollamaHost.Host, ollamaHost.Port)
	}

	client := ollama.NewClient(
		&url.URL{
			Scheme: ollamaHost.Scheme,
			Host:   config.Url,
		},
		http.DefaultClient,
	)
	return &CopilotClient{
		ollama:      client,
		model:       config.ModelName,
		temperature: config.Temperature,
		topP:        config.TopP,
	}
}
