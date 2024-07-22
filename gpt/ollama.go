package gpt

import (
	"context"

	rpc "github.com/cligpt/shdrive/drive/rpc"
)

const (
	OllamaName = "llama3"
)

func OllamaChat(_ context.Context, cfg *Config, req *rpc.ChatRequest) (*rpc.ChatReply, error) {
	// TBD: FIXME
	return &rpc.ChatReply{
		Model: &rpc.ChatModel{
			Name: OllamaName,
			Id:   "",
			Key:  "",
		},
		CreatedAt: "2023-08-04T08:52:19.385406455-07:00",
		Message: &rpc.ChatMessage{
			Role:    "user",
			Content: "content",
		},
		Done: true,
	}, nil
}
