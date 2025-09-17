package copilot

import (
	"context"
	"fmt"
	"testing"

	"github.com/lw396/ChatCopilot/internal/repository/gorm"
)

func TestChat(t *testing.T) {
	client, err := NewClient(&gorm.CopilotConfig{
		ModelName:   "qwen2",
		Temperature: 0.1,
		TopP:        0.1,
	})
	if err != nil {
		t.Error("erorr:", err)
	}

	messages := []Message{{
		Role:    RoleUser,
		Content: "你好",
	}}
	ch := make(chan interface{})
	err = client.Chat(context.Background(), messages, ch)
	if err != nil {
		t.Error("erorr:", err)
	}

	for val := range ch {
		fmt.Println(val)
	}

}
