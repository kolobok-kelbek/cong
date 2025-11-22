# Cong

## Golang library for convenient configuration management

![CI](https://github.com/kolobok-kelbek/cong/actions/workflows/test.yml/badge.svg)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.25-61CFDD.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolobok-kelbek/cong)](https://goreportcard.com/report/github.com/kolobok-kelbek/cong)
[![Go Reference](https://pkg.go.dev/badge/github.com/kolobok-kelbek/cong.svg)](https://pkg.go.dev/github.com/kolobok-kelbek/cong)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

The library inside uses [viper](https://github.com/spf13/viper).

Provides the ability to use a pre-configured viper through a simplified initialization interface. Additionally maps
CamelCase structure fields to snake_case environment variables (e.g., `ServerName` -> `SERVER_NAME` instead of `SERVERNAME`).

## Features
- One-line loader helpers for files, directories, embed.FS, or pure environment variables.
- Automatic env binding with snake_case + project prefix (good for 12-factor apps).
- Supports common formats (JSON, YAML, TOML, dotenv, HCL, ini, …) via viper.
- Generic-friendly: `loader := cong.NewLoader[MyConfig]()`.

## Installation

```bash
go get github.com/kolobok-kelbek/cong
```

## Loader methods

1. `Load(projectName string, ext ConfigExtension, configPaths ...string)` — read single file by name+extension from provided or default paths (".", "./config", "./static").
2. `LoadFromDir(projectName string, path string, ext ConfigExtension)` — merge all files with the extension from a directory tree.
3. `LoadFromEmbedFS(projectName string, dir embed.FS, ext ConfigExtension)` — merge all files with the extension from embed.FS.
4. `LoadFromEmbedFSByPath(projectName string, dir embed.FS, path string, ext ConfigExtension)` — same as above but scoped to a path.
5. `LoadFromEnv(projectName string)` — load only from environment variables (12-factor friendly).

## Examples

Run any of the ready-to-use examples:

- Basic file load: `go run ./examples/basic-file` (reads `examples/basic-file/config/config.yaml` via `Load`).
- Env-only (12-factor): `go run ./examples/env-only` (`LoadFromEnv`, no files).
- Embedded configs: `go run ./examples/embed` (uses `embed.FS` + `LoadFromEmbedFS`).

Inline YAML + env override example:

```yaml
serverName: HelloWorld
port: 80
timeout: 20
```

```golang
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kolobok-kelbek/cong"
)

type Config struct {
	ServerName string `mapstructure:"serverName"`
	Port       int    `mapstructure:"port"`
	Timeout    int    `mapstructure:"timeout"`
}

func main() {
	_ = os.Setenv("HELLO_SERVER_NAME", "PAM-PAM") // overrides serverName

	loader := cong.NewLoader[Config]()
	cfg, err := loader.Load("config", cong.YamlExt)
	if err != nil {
		panic(err)
	}

	out, _ := json.Marshal(cfg)
	fmt.Println(string(out))
}
```

### Twelve-Factor / env-only usage

If you don't want to ship any config files (12-factor style), you can populate the config only from environment variables:

```golang
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kolobok-kelbek/cong"
)

type Config struct {
	ServerName string
	Port       int
	Timeout    int
}

func main() {
	_ = os.Setenv("HELLO_SERVER_NAME", "pam-pam")
	_ = os.Setenv("HELLO_PORT", "8080")
	_ = os.Setenv("HELLO_TIMEOUT", "30")

	loader := cong.NewLoader[Config]()
	cfg, err := loader.LoadFromEnv("hello")
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(cfg)
	fmt.Println(string(b))
}
```

> More examples you can find in [loader_test.go](loader_test.go) and the [examples](examples) directory.

## Tasks

If you use [Task](https://taskfile.dev):
- `task lint` / `task lint-fix` – run golangci-lint (fix where supported).
- `task test` – run the test suite.
