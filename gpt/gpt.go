package gpt

import (
	"context"

	"github.com/hashicorp/go-hclog"

	"github.com/cligpt/shdrive/config"
	rpc "github.com/cligpt/shdrive/drive/rpc"
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
	// TBD: FIXME

	return nil, nil
}
