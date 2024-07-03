package copilot

import (
	"net/http"
	"net/url"

	ollama "github.com/ollama/ollama/api"
)

func NewClient(base *url.URL, http *http.Client) *ollama.Client {
	return ollama.NewClient(base, http)
}

func ClientFromEnvironment() (*ollama.Client, error) {
	return ollama.ClientFromEnvironment()
}
