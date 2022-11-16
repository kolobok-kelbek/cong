package cong

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
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

func (loader *Loader[T]) Load(projectName string, ext ConfigExtension, configPaths ...string) (*T, error) {
	loader.viper.SetConfigName(projectName)
	loader.viper.SetEnvPrefix(projectName)

	loader.viper.SetConfigType(ext.String())

	loader.loadConfigPaths(configPaths)

	err := loader.viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := new(T)
	err = loader.viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (loader *Loader[T]) LoadDir(path string, ext ConfigExtension) (*T, error) {
	configsPaths, err := loader.findConfigFilesInDir(path, ext)
	if err != nil {
		return nil, err
	}

	err = loader.loadConfigFilesByPaths(configsPaths, ext)
	if err != nil {
		return nil, err
	}

	config := new(T)
	err = loader.viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (loader *Loader[T]) loadConfigFilesByPaths(configsPaths []string, ext ConfigExtension) error {
	for _, path := range configsPaths {
		dir, file := filepath.Split(path)
		configName := file[:len(file)-len(filepath.Ext(file))]
		loader.viper.SetConfigName(configName)
		loader.viper.SetConfigType(ext.String())
		loader.viper.AddConfigPath(dir)
		err := loader.viper.MergeInConfig()
		if err != nil {
			return err
		}
	}

	return nil
}

func (loader *Loader[T]) findConfigFilesInDir(path string, ext ConfigExtension) ([]string, error) {
	var configsPaths []string

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(absolutePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == "."+ext.String() {
			configsPaths = append(configsPaths, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return configsPaths, nil
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
