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

	cfg, err := loader.Load("config", cong.YamlExt, "./examples/basic-file/config")
	if err != nil {
		panic(err)
	}

	out, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(out))
}
