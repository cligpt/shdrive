package drive

import (
	"context"
	"github.com/cligpt/shdrive/config"
	"github.com/cligpt/shdrive/etcd"
	"github.com/cligpt/shdrive/gpt"
	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
)

type Drive interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Logger hclog.Logger
	Config config.Config
	Etcd   etcd.Etcd
	Gpt    gpt.Gpt
	Http   string
	Rpc    string
}

type drive struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Drive {
	return &drive{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (d *drive) Init(ctx context.Context) error {
	if err := d.cfg.Etcd.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init etcd")
	}

	if err := d.cfg.Gpt.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init gpt")
	}

	return nil
}

func (d *drive) Deinit(ctx context.Context) error {
	_ = d.cfg.Gpt.Deinit(ctx)
	_ = d.cfg.Etcd.Deinit(ctx)

	return nil
}

func (d *drive) Run(_ context.Context) error {
	// TBD: FIXME
	return nil
}
