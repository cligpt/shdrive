package drive

import (
	"context"
	"math"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/cligpt/shdrive/config"
	rpc "github.com/cligpt/shdrive/drive/rpc"
	"github.com/cligpt/shdrive/etcd"
	"github.com/cligpt/shdrive/gpt"
)

const (
	httpTimeout = 30 * time.Second
)

type Drive interface {
	Init(context.Context) error
	Deinit(context.Context) error
	RunHttp(context.Context) error
	RunRpc(context.Context) error
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
	cfg     *Config
	srvHttp *http.Server
	srvRpc  *grpc.Server
	rpc.UnimplementedRpcProtoServer
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", d.getRoot)
	mux.HandleFunc("/hello", d.getHello)
	d.srvHttp = &http.Server{
		Addr:              d.cfg.Http,
		Handler:           mux,
		ReadTimeout:       httpTimeout,
		ReadHeaderTimeout: httpTimeout,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}

	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}
	d.srvRpc = grpc.NewServer(options...)
	rpc.RegisterRpcProtoServer(d.srvRpc, d)

	return nil
}

func (d *drive) Deinit(ctx context.Context) error {
	d.srvRpc.Stop()
	_ = d.srvHttp.Close()
	_ = d.cfg.Gpt.Deinit(ctx)
	_ = d.cfg.Etcd.Deinit(ctx)

	return nil
}

func (d *drive) RunHttp(ctx context.Context) error {
	return d.srvHttp.ListenAndServe()
}

func (d *drive) RunRpc(_ context.Context) error {
	lis, err := net.Listen("tcp", d.cfg.Rpc)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	return d.srvRpc.Serve(lis)
}

func (d *drive) SendChat(_ context.Context, in *rpc.ChatRequest) (*rpc.ChatReply, error) {
	// TBD: FIXME
	return &rpc.ChatReply{
		Model: &rpc.ChatModel{
			Name: "llama3",
			Id:   "",
			Key:  "",
		},
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
