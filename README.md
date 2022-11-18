# Cong

## Golang library for convenient configuration management

![Go Version](https://img.shields.io/badge/go%20version-%3E=1.19-61CFDD.svg)

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