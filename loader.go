package cong

import (
	"fmt"
	"github.com/spf13/viper"
)

var defaultConfigPaths = []string{
	".",
	"./config",
	"./static",
}

type Loader[T any] struct {
	viper *viper.Viper
}

func NewLoader[T any]() *Loader[T] {
	return &Loader[T]{
		viper: viper.New(),
	}
}

func (loader *Loader[T]) Load(projectName string, ext ConfigExtension, configPaths ...string) *T {
	config := new(T)

	loader.viper.SetConfigName(projectName)
	loader.viper.SetEnvPrefix(projectName)

	loader.viper.SetConfigType(ext.String())

	loader.loadConfigPaths(configPaths)

	err := loader.viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = loader.viper.Unmarshal(config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return config
}

func (loader *Loader[T]) loadConfigPaths(configPaths []string) {
	var paths []string
	if len(configPaths) != 0 {
		paths = configPaths
	} else {
		paths = defaultConfigPaths
	}

	for _, path := range paths {
		loader.viper.AddConfigPath(path)
	}
}
