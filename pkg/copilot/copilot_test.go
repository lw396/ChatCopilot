package copilot

import (
	"context"
	"testing"

	ollama "github.com/ollama/ollama/api"
)

func TestChat(t *testing.T) {
	ollamaclient, err := ClientFromEnvironment()
	if err != nil {
		return
	}

	client := &CopilotClient{
		ollama: ollamaclient,
		model:  "qwen2",
		stream: true,
	}

	messages := []ollama.Message{
		{
			Role:    "user",
			Content: "你好",
		},
	}

	if err := client.Chat(context.Background(), messages); err != nil {
		t.Error("erorr:", err)
	}
}
