# shdrive

[![Build Status](https://github.com/cligpt/shdrive/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/cligpt/shdrive/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/cligpt/shdrive/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/cligpt/shdrive)
[![Go Report Card](https://goreportcard.com/badge/github.com/cligpt/shdrive)](https://goreportcard.com/report/github.com/cligpt/shdrive)
[![License](https://img.shields.io/github/license/cligpt/shdrive.svg)](https://github.com/cligpt/shdrive/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/cligpt/shdrive.svg)](https://github.com/cligpt/shdrive/tags)



## Introduction

*shdrive* is the server of [shai](https://github.com/cligpt/shai) written in Go.



## Prerequisites

- Go >= 1.22.0



## Build

```bash
version=latest
make build
```



## Usage

```
shai server

Usage:
  shdrive [flags]

Flags:
  -f, --config-file string   config file (default "$HOME/.shai/shdrive.yml")
  -h, --help                 help for shdrive
  -t, --listen-http string   listen http (default ":69090")
  -r, --listen-rpc string    listen rpc (default ":65090")
  -l, --log-level string     log level (DEBUG|INFO|WARN|ERROR) (default "WRAN")
  -v, --version              version for shdrive
```



## License

Project License can be found [here](LICENSE).



## Reference

- [warp](https://www.warp.dev/)
