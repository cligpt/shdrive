package drive

import (
	"context"
	"math"
	"net"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	"github.com/cligpt/shdrive/config"
	rpc "github.com/cligpt/shdrive/drive/rpc"
)

type Up interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, string) error
}

type UpConfig struct {
	Logger hclog.Logger
	Config config.Config
}

type up struct {
	cfg *Config
	srv *grpc.Server
	rpc.UnimplementedUpProtoServer
}

func UpNew(_ context.Context, cfg *Config) Up {
	return &up{
		cfg: cfg,
	}
}

func UpDefaultConfig() *UpConfig {
	return &UpConfig{}
}

func (u *up) Init(ctx context.Context) error {
	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}

	u.srv = grpc.NewServer(options...)

	rpc.RegisterUpProtoServer(u.srv, u)

	return nil
}

func (u *up) Deinit(ctx context.Context) error {
	u.srv.Stop()

	return nil
}

func (u *up) Run(_ context.Context, url string) error {
	lis, _ := net.Listen("tcp", url)

	return u.srv.Serve(lis)
}

func (a *ai) SendQuery(ctx context.Context, in *rpc.QueryRequest) (*rpc.QueryReply, error) {
	// TBD: FIXME
	return &rpc.QueryReply{
		Version: "v0.1.0",
		Url:     "https://github.com/cligpt/shai/releases/download/v0.1.0/shai_0.1.0_linux_amd64.tar.gz",
		User:    "",
		Pass:    "",
	}, nil
}
