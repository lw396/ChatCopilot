package copilot

import (
	"context"
	"fmt"
	"testing"

	"github.com/lw396/ChatCopilot/internal/repository/gorm"
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
	ch := make(chan interface{})
	err := client.Chat(context.Background(), messages, ch)
	if err != nil {
		t.Error("erorr:", err)
	}

	for val := range ch {
		fmt.Println(val)
	}

}
