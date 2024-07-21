package gpt

import (
	"context"

	rpc "github.com/cligpt/shdrive/drive/rpc"
)

const (
	ollamaName = "llama3"
)

func ollamaChat(_ context.Context, _ *rpc.ChatRequest) (*rpc.ChatReply, error) {
	// TBD: FIXME
	return &rpc.ChatReply{
		Model: &rpc.ChatModel{
			Name: ollamaName,
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
