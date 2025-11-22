package main

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/kolobok-kelbek/cong"
)

//go:embed config/*.yaml
var embeddedConfigs embed.FS

type App struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
}

type Server struct {
	Port    int `mapstructure:"port"`
	Timeout int `mapstructure:"timeout"`
}

type Config struct {
	App    App    `mapstructure:"app"`
	Server Server `mapstructure:"server"`
}

func main() {
	loader := cong.NewLoader[Config]()

	cfg, err := loader.LoadFromEmbedFS("embedded", embeddedConfigs, cong.YamlExt)
	if err != nil {
		panic(err)
	}

	out, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(out))
}
