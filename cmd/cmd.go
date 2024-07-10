package cmd

import (
	"context"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/hashicorp/go-hclog"
	"github.com/pkg/errors"

	"github.com/cligpt/shdrive/config"
	"github.com/cligpt/shdrive/drive"
)

const (
	driveName = "shdrive"
)

var (
	app      = kingpin.New(driveName, "shai server").Version(config.Version + "-build-" + config.Build)
	logLevel = app.Flag("log-level", "Log level (DEBUG|INFO|WARN|ERROR)").Default("WARN").String()
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

	d, err := initDrive(ctx, logger, c)
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

func initDrive(ctx context.Context, logger hclog.Logger, _ *config.Config) (drive.Drive, error) {
	c := drive.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Logger = logger

	return drive.New(ctx, c), nil
}

func runDrive(ctx context.Context, _ hclog.Logger, _drive drive.Drive) error {
	if err := _drive.Init(ctx); err != nil {
		return errors.New("failed to init")
	}

	defer func(_drive drive.Drive, ctx context.Context) {
		_ = _drive.Deinit(ctx)
	}(_drive, ctx)

	if err := _drive.Run(ctx); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}
