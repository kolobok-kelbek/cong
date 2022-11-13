package main

import (
	"encoding/json"
	"fmt"
	"github.com/kolobok-kelbek/cong"
)

type Config struct {
	Port    int
	Timeout int
}

func main() {
	loader := cong.NewLoader[Config]()
	data := loader.Load("example", cong.YamlExt)

	marshal, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("fatal error marshal config data: %w", err))
	}

	fmt.Println(string(marshal))
}
