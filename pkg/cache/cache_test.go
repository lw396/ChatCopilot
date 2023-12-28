package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDefaultCache(t *testing.T) {
	val := []int32{1, 2, 3}
	store := DefaultStore()
	key := "TEST_VAL"

	if err := store.Set(context.Background(), key, val, 1*time.Second); err != nil {
		t.Error(err)
		return
	}

	var result []int32
	found, err := store.Get(context.Background(), key, &result)
	if err != nil {
		t.Error(err)
		return
	}
	if !found {
		t.Error("can not found value")
		return
	}
	fmt.Println(result)
}
