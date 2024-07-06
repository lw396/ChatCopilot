package copilot

import (
	"context"
	"testing"

	"github.com/lw396/WeComCopilot/internal/repository/gorm"
	ollama "github.com/ollama/ollama/api"
)

func TestChat(t *testing.T) {
	client := NewClient(&gorm.CopilotConfig{
		ModelName:   "qwen2",
		Temperature: 0.1,
		TopP:        0.1,
	})

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
