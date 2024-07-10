package gpt

import (
	"context"

	"github.com/hashicorp/go-hclog"

	"github.com/cligpt/shdrive/config"
)

type Gpt interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
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

func (g *gpt) Run(_ context.Context) error {
	return nil
}
