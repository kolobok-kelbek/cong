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
	// Only env vars are used; no config files needed.
	_ = os.Setenv("HELLO_SERVER_NAME", "pam-pam")
	_ = os.Setenv("HELLO_PORT", "8080")
	_ = os.Setenv("HELLO_TIMEOUT", "45")

	loader := cong.NewLoader[Config]()

	cfg, err := loader.LoadFromEnv("hello")
	if err != nil {
		panic(err)
	}

	out, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(out))
}
