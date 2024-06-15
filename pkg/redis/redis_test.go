package redis

import (
	"context"
	"testing"
)

func TestSAdd(t *testing.T) {
	client, err := NewRedisClient()
	if err != nil {
		t.Fatal(err)
	}
	data := []struct {
		Name string
		Age  int
	}{
		{
			Name: "Tom",
			Age:  2,
		},
		{
			Name: "Jerry",
			Age:  4,
		},
	}

	for _, d := range data {
		if err := client.SAdd(context.Background(), "TEST_MEMNERS", d); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSRem(t *testing.T) {
	client, err := NewRedisClient()
	if err != nil {
		t.Fatal(err)
	}
	data := []struct {
		Name string
		Age  int
	}{
		{
			Name: "Tom",
			Age:  2,
		},
		{
			Name: "Jerry",
			Age:  4,
		},
	}
	for _, d := range data {
		if err := client.SRem(context.Background(), "TEST_MEMNERS", &d); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSMembers(t *testing.T) {
	client, err := NewRedisClient()
	if err != nil {
		t.Fatal(err)
	}

	data := []*struct {
		Name string
		Age  int
	}{}
	found, err := client.SMembers(context.Background(), "TEST_MEMNERS", &data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(found, data[0].Name)
}

func NewRedisClient() (RedisClient, error) {
	return NewClient(
		WithAddress("127.0.0.1", 6379),
		WithAuth("", "secret"),
		WithDB(0),
	)
}
