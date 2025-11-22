# Cong

## Golang library for convenient configuration management

![CI](https://github.com/kolobok-kelbek/cong/actions/workflows/test.yml/badge.svg)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.25-61CFDD.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kolobok-kelbek/cong)](https://goreportcard.com/report/github.com/kolobok-kelbek/cong)
[![Go Reference](https://pkg.go.dev/badge/github.com/kolobok-kelbek/cong.svg)](https://pkg.go.dev/github.com/kolobok-kelbek/cong)

The library inside uses [viper](https://github.com/spf13/viper)

Provides the ability to use a pre-configured viper through a simplified initialization interface.\
Additionally able to map CamelCase structure fields to SnakeCase environment variables.
**Example**: field `ServerName` is going to be mapped with environment variable `SERVER_NAME`,in the original viper work
differently - field `ServerName` is going to be mapped with environment variable `SERVERNAME`.


The interface has 4 loader methods, the choice depends on how you plan to load configuration files:

1. `Load(projectName string, ext ConfigExtension, configPaths ...string) (*T, error)` - downloads a single file by 
    project name and extension. The file is searched in directories by the passed paths or in default directories (".", 
    "./config", "./static"). The directory search is sequential, only the first found file will be loaded. Loader reads 
    the config file and unmarshal it into an generic type object, the object is returned for further use.
2. `LoadFromDir(projectName string, path string, ext ConfigExtension) (*T, error)` - all files from the directory with 
    the passed extension are loaded and unmarshal.
3. `LoadFromEmbedFS(projectName string, dir embed.FS, ext ConfigExtension) (*T, error)`-all files from the embed.FS 
    with the passed extension are loaded and unmarshal.
4. `LoadFromEmbedFSByPath(projectName string, dir embed.FS, path string, ext ConfigExtension) (*T, error)`-all files by 
    path from the embed.FS with the passed extension are loaded and unmarshal.
5. `LoadFromEnv(projectName string) (*T, error)` - fills config only from environment variables, which is handy for 
    Twelve-Factor applications where no config files should be mounted.

### Examples

#### JSON

```json
{
  "port": 80,
  "timeout": "10s"
}
```

```golang
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kolobok-kelbek/cong"
)

type Config struct {
	Port    int
	Timeout string
}

func main() {
	loader := cong.NewLoader[Config]()
	data, err := loader.Load("config", cong.JsonExt)
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("fatal error marshal config data: %w", err))
	}

fmt.Println(string(marshal))
}

```

### Twelve-Factor / env-only usage

If you don't want to ship any config files (12-factor style), you can populate the config only from environment variables:

```golang
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kolobok-kelbek/cong"
	"os"
)

type Config struct {
	ServerName string
	Port       int
	Timeout    int
}

func main() {
	os.Setenv("HELLO_SERVER_NAME", "pam-pam")
	os.Setenv("HELLO_PORT", "8080")
	os.Setenv("HELLO_TIMEOUT", "30")

	loader := cong.NewLoader[Config]()
	cfg, err := loader.LoadFromEnv("hello")
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(cfg)
	fmt.Println(string(b))
}
```

#### YAML

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
	"github.com/kolobok-kelbek/cong"
)

type Config struct {
	ServerName string `mapstructure:"serverName"`
	Port       int    `mapstructure:"port"`
	Timeout    string `mapstructure:"timeout"`
}

func main() {
	loader := cong.NewLoader[Config]()
	data, err := loader.Load("config", cong.YamlExt)
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("fatal error marshal config data: %w", err))
	}

	fmt.Println(string(marshal))
}

```

#### YAML and binding environment variables

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
	"github.com/kolobok-kelbek/cong"
	"os"
)

type Config struct {
	ServerName string `mapstructure:"serverName"`
	Port       int    `mapstructure:"port"`
	Timeout    string `mapstructure:"timeout"`
}

func main() {
	os.Setenv("CONFIG_SERVER_NAME", "pam-pam")

	loader := cong.NewLoader[Config]()
	data, err := loader.Load("config", cong.YamlExt)
	if err != nil {
		panic(err)
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("fatal error marshal config data: %w", err))
	}

	fmt.Println(string(marshal))
}

```

> More examples you can find in [loader_test.go](loader_test.go)
