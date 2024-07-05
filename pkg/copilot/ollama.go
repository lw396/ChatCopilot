package copilot

import (
	"net/http"
	"net/url"

	ollama "github.com/ollama/ollama/api"
)

func NewClient(base *url.URL, http *http.Client) *CopilotClient {
	return &CopilotClient{
		ollama: ollama.NewClient(base, http),
		model:  "qwen",
		stream: true,
	}
}

func ClientFromEnvironment() (*ollama.Client, error) {
	return ollama.ClientFromEnvironment()
}
