package etcd

import (
	"context"
	"github.com/hashicorp/go-hclog"

	"github.com/cligpt/shdrive/config"
)

type Etcd interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Logger hclog.Logger
	Config config.Config
}

type etcd struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Etcd {
	return &etcd{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (e *etcd) Init(ctx context.Context) error {
	return nil
}

func (e *etcd) Deinit(_ context.Context) error {
	return nil
}

func (e *etcd) Run(_ context.Context) error {
	return nil
}
