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

type Ai interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, string) error
}

type AiConfig struct {
	Logger hclog.Logger
	Config config.Config
}

type ai struct {
	cfg *Config
	srv *grpc.Server
	rpc.UnimplementedAiProtoServer
}

func AiNew(_ context.Context, cfg *Config) Ai {
	return &ai{
		cfg: cfg,
	}
}

func AiDefaultConfig() *AiConfig {
	return &AiConfig{}
}

func (a *ai) Init(ctx context.Context) error {
	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}

	a.srv = grpc.NewServer(options...)

	rpc.RegisterAiProtoServer(a.srv, a)

	return nil
}

func (a *ai) Deinit(ctx context.Context) error {
	a.srv.Stop()

	return nil
}

func (a *ai) Run(_ context.Context, url string) error {
	lis, _ := net.Listen("tcp", url)

	return a.srv.Serve(lis)
}

func (a *ai) SendChat(ctx context.Context, in *rpc.ChatRequest) (*rpc.ChatReply, error) {
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
