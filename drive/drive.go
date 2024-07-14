package drive

import (
	"context"
	"math"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/cligpt/shdrive/config"
	rpc "github.com/cligpt/shdrive/drive/rpc"
	"github.com/cligpt/shdrive/etcd"
	"github.com/cligpt/shdrive/gpt"
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
	srv *grpc.Server
	rpc.UnimplementedAiProtoServer
	rpc.UnimplementedUpProtoServer
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

	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}

	d.srv = grpc.NewServer(options...)

	rpc.RegisterAiProtoServer(d.srv, d)
	rpc.RegisterUpProtoServer(d.srv, d)

	return nil
}

func (d *drive) Deinit(ctx context.Context) error {
	d.srv.Stop()

	_ = d.cfg.Gpt.Deinit(ctx)
	_ = d.cfg.Etcd.Deinit(ctx)

	return nil
}

func (d *drive) Run(_ context.Context) error {
	// TBD: FIXME
	return nil
}
