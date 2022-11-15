# Cong

## Golang library for convenient configuration management

![Go Version](https://img.shields.io/badge/go%20version-%3E=1.19-61CFDD.svg)

# example

```
type Config struct {
	Port    int
	Timeout int
}

func main() {
	loader := cong.NewLoader[Config]()
	data := loader.Load("example", cong.JsonExt)

	marshal, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("fatal error marshal config data: %w", err))
	}

	fmt.Println(string(marshal))
}
```