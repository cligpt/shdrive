package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kingpin/v2"
	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/cligpt/shdrive/config"
	"github.com/cligpt/shdrive/drive"
	"github.com/cligpt/shdrive/etcd"
	"github.com/cligpt/shdrive/gpt"
)

const (
	driveName  = "shdrive"
	routineNum = -1
)

var (
	app      = kingpin.New(driveName, "shai server").Version(config.Version + "-build-" + config.Build)
	logLevel = app.Flag("log-level", "Log level (DEBUG|INFO|WARN|ERROR)").Short('l').Default("WARN").String()
)

func Run(ctx context.Context) error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	logger, err := initLogger(ctx, *logLevel)
	if err != nil {
		return errors.Wrap(err, "failed to init logger")
	}

	c, err := initConfig(ctx, logger)
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	e, err := initEtcd(ctx, logger, c)
	if err != nil {
		return errors.Wrap(err, "failed to init etcd")
	}

	g, err := initGpt(ctx, logger, c)
	if err != nil {
		return errors.Wrap(err, "failed to init gpt")
	}

	d, err := initDrive(ctx, logger, c, e, g)
	if err != nil {
		return errors.Wrap(err, "failed to init drive")
	}

	if err := runDrive(ctx, logger, d); err != nil {
		return errors.Wrap(err, "failed to run drive")
	}

	return nil
}

func initLogger(_ context.Context, level string) (hclog.Logger, error) {
	return hclog.New(&hclog.LoggerOptions{
		Name:  driveName,
		Level: hclog.LevelFromString(level),
	}), nil
}

func initConfig(_ context.Context, _ hclog.Logger) (*config.Config, error) {
	c := config.New()
	return c, nil
}

func initEtcd(ctx context.Context, logger hclog.Logger, cfg *config.Config) (etcd.Etcd, error) {
	c := etcd.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Logger = logger
	c.Config = *cfg

	return etcd.New(ctx, c), nil
}

func initGpt(ctx context.Context, logger hclog.Logger, cfg *config.Config) (gpt.Gpt, error) {
	c := gpt.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Logger = logger
	c.Config = *cfg

	return gpt.New(ctx, c), nil
}

func initDrive(ctx context.Context, logger hclog.Logger, cfg *config.Config, _etcd etcd.Etcd, _gpt gpt.Gpt) (drive.Drive, error) {
	c := drive.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Logger = logger
	c.Config = *cfg
	c.Etcd = _etcd
	c.Gpt = _gpt

	return drive.New(ctx, c), nil
}

func runDrive(ctx context.Context, _ hclog.Logger, _drive drive.Drive) error {
	if err := _drive.Init(ctx); err != nil {
		return errors.New("failed to init")
	}

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(routineNum)

	g.Go(func() error {
		if err := _drive.Run(ctx); err != nil {
			return errors.Wrap(err, "failed to run")
		}
		return nil
	})

	s := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be caught, so don't need add it
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	g.Go(func() error {
		<-s
		_ = _drive.Deinit(ctx)
		return nil
	})

	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "failed to wait")
	}

	return nil
}
