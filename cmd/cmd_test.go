package cmd

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/cligpt/shdrive/config"
)

func testInitConfig() *config.Config {
	cfg := config.New()

	fi, _ := os.Open("../test/config/config.yml")

	defer func() {
		_ = fi.Close()
	}()

	buf, _ := io.ReadAll(fi)
	_ = yaml.Unmarshal(buf, cfg)

	return cfg
}

func TestInitLogger(t *testing.T) {
	_, err := initLogger(context.Background(), "WARN")
	assert.Equal(t, nil, err)
}

func TestInitEtcd(t *testing.T) {
	cfg := testInitConfig()
	_logger, _ := initLogger(context.Background(), "WARN")

	_, err := initEtcd(context.Background(), _logger, cfg)
	assert.Equal(t, nil, err)
}

func TestInitGpt(t *testing.T) {
	cfg := testInitConfig()
	_logger, _ := initLogger(context.Background(), "WARN")

	_, err := initGpt(context.Background(), _logger, cfg)
	assert.Equal(t, nil, err)
}

func TestInitDrive(t *testing.T) {
	cfg := testInitConfig()
	_logger, _ := initLogger(context.Background(), "WARN")
	_etcd, _ := initEtcd(context.Background(), _logger, cfg)
	_gpt, _ := initGpt(context.Background(), _logger, cfg)
	_http := ":68080"
	_rpc := ":65050"

	_, err := initDrive(context.Background(), _logger, cfg, _etcd, _gpt, _http, _rpc)
	assert.Equal(t, nil, err)
}
