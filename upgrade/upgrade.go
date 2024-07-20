package upgrade

import (
	"context"

	"github.com/hashicorp/go-hclog"

	"github.com/cligpt/shdrive/config"
	rpc "github.com/cligpt/shdrive/drive/rpc"
)

type Upgrade interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, *rpc.QueryRequest) (*rpc.QueryReply, error)
}

type Config struct {
	Logger hclog.Logger
	Config config.Config
}

type upgrade struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Upgrade {
	return &upgrade{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (u *upgrade) Init(_ context.Context) error {
	return nil
}

func (u *upgrade) Deinit(_ context.Context) error {
	return nil
}

func (u *upgrade) Run(ctx context.Context, req *rpc.QueryRequest) (*rpc.QueryReply, error) {
	// TBD: FIXME
	return &rpc.QueryReply{
		Version: "v0.1.0",
		Url:     "https://github.com/cligpt/shai/releases/download/v0.1.0/shai_0.1.0_linux_amd64.tar.gz",
		User:    "",
		Pass:    "",
	}, nil
}
