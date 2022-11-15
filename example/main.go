package main

import (
	"encoding/json"
	"fmt"
	"github.com/kolobok-kelbek/cong"
	"os"
)

type Config struct {
	Port    int
	Timeout int
}

func main() {
	loader := cong.NewLoader[Config]()
	data, err := loader.Load("example", cong.YamlExt)
	if err != nil {
		panic(fmt.Errorf("fatal error parce config: %w", err))
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("fatal error marshal config data: %w", err))
	}

	fmt.Println(string(marshal))
	fmt.Println(os.Getwd())

	loader = cong.NewLoader[Config]()
	data, err = loader.LoadDir("./configDir", cong.YamlExt)
	if err != nil {
		panic(fmt.Errorf("fatal error parce config: %w", err))
	}

	marshal, err = json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("fatal error marshal config data: %w", err))
	}

	fmt.Println(string(marshal))
}
