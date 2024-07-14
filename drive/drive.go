package drive

import (
	"context"
	rpc "github.com/cligpt/shdrive/drive/rpc"
	"google.golang.org/grpc"
	"math"
	"net"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"

	"github.com/cligpt/shdrive/config"
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
	cfg      *Config
	aiServer *grpc.Server
	rpc.UnimplementedAiProtoServer
	upServer *grpc.Server
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

	d.aiServer = grpc.NewServer(options...)
	rpc.RegisterAiProtoServer(d.aiServer, d)

	d.upServer = grpc.NewServer(options...)
	rpc.RegisterUpProtoServer(d.upServer, d)

	return nil
}

func (d *drive) Deinit(ctx context.Context) error {
	d.upServer.Stop()
	d.aiServer.Stop()

	_ = d.cfg.Gpt.Deinit(ctx)
	_ = d.cfg.Etcd.Deinit(ctx)

	return nil
}

func (d *drive) Run(_ context.Context) error {
	lis, _ := net.Listen("tcp", d.cfg.Rpc)

	if err := d.aiServer.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to run ai server")
	}

	if err := d.upServer.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to run up server")
	}

	// TBD: FIXME
	// Run http server

	return nil
}

func (d *drive) SendChat(_ context.Context, in *rpc.ChatRequest) (*rpc.ChatReply, error) {
	// TBD: FIXME
	return &rpc.ChatReply{
		Model:     "llama3",
		CreatedAt: "2023-08-04T08:52:19.385406455-07:00",
		Message: &rpc.ChatMessage{
			Role:    "user",
			Content: "content",
		},
		Done: true,
	}, nil
}

func (d *drive) SendQuery(_ context.Context, in *rpc.QueryRequest) (*rpc.QueryReply, error) {
	// TBD: FIXME
	return &rpc.QueryReply{
		Version: "v0.1.0",
		Url:     "https://github.com/cligpt/shai/releases/download/v0.1.0/shai_0.1.0_linux_amd64.tar.gz",
		User:    "",
		Pass:    "",
	}, nil
}
