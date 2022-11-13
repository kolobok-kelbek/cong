# Cong

## Golang library for convenient configuration management

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