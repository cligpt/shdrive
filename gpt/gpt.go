package gpt

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"

	"github.com/cligpt/shdrive/config"
	rpc "github.com/cligpt/shdrive/drive/rpc"
)

var (
	gptModel = map[string]func(context.Context, *rpc.ChatRequest) (*rpc.ChatReply, error){
		ollamaName: ollamaChat,
	}
)

type Gpt interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, *rpc.ChatRequest) (*rpc.ChatReply, error)
}

type Config struct {
	Logger hclog.Logger
	Config config.Config
}

type gpt struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Gpt {
	return &gpt{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (g *gpt) Init(_ context.Context) error {
	return nil
}

func (g *gpt) Deinit(_ context.Context) error {
	return nil
}

func (g *gpt) Run(ctx context.Context, req *rpc.ChatRequest) (*rpc.ChatReply, error) {
	var err error
	var rep *rpc.ChatReply

	if _, ok := gptModel[req.Model.Name]; !ok {
		return nil, errors.New("invalid model name")
	}

	rep, err = gptModel[req.Model.Name](ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to run model")
	}

	return rep, nil
}
