package cong

import (
	"embed"
	"github.com/spf13/viper"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"unicode"
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
	config := new(T)

	loader.setDefaultSettings(projectName)

	loader.bindSnakeCaseParams(config)

	loader.viper.SetConfigName(projectName)
	loader.viper.SetConfigType(ext.String())

	loader.loadConfigPaths(configPaths)

	err := loader.viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = loader.viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (loader *Loader[T]) LoadFromDir(projectName string, path string, ext ConfigExtension) (*T, error) {
	config := new(T)

	loader.setDefaultSettings(projectName)

	loader.bindSnakeCaseParams(config)

	configsPaths, err := loader.findConfigFilesInDir(path, ext)
	if err != nil {
		return nil, err
	}

	err = loader.loadConfigFilesByPaths(configsPaths, ext)
	if err != nil {
		return nil, err
	}

	err = loader.viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (loader *Loader[T]) LoadFromEmbedFS(projectName string, dir embed.FS, ext ConfigExtension) (*T, error) {
	config := new(T)

	loader.bindSnakeCaseParams(config)

	loader.setDefaultSettings(projectName)

	configsPaths, err := loader.findConfigFilesInEmbedFS(".", dir, ext)
	if err != nil {
		return nil, err
	}

	err = loader.loadConfigFilesByPaths(configsPaths, ext)
	if err != nil {
		return nil, err
	}

	err = loader.viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (loader *Loader[T]) LoadFromEmbedFSByPath(
	projectName string,
	dir embed.FS,
	path string,
	ext ConfigExtension,
) (*T, error) {
	config := new(T)

	loader.bindSnakeCaseParams(config)

	loader.setDefaultSettings(projectName)

	configsPaths, err := loader.findConfigFilesInEmbedFS(path, dir, ext)
	if err != nil {
		return nil, err
	}

	err = loader.loadConfigFilesByPaths(configsPaths, ext)
	if err != nil {
		return nil, err
	}

	err = loader.viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (loader *Loader[T]) setDefaultSettings(projectName string) {
	loader.viper.AutomaticEnv()
	loader.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	loader.viper.SetEnvPrefix(projectName)
}

func (loader *Loader[T]) bindSnakeCaseParams(config *T) {
	refVal := reflect.ValueOf(config).Elem()

	for i := 0; i < refVal.NumField(); i++ {
		name := refVal.Type().Field(i).Name
		fieldName := strings.ToUpper(name)
		paramName := strings.ToUpper(loader.toSnakeCase(name))

		if fieldName == paramName {
			continue
		}

		loader.viper.SetEnvKeyReplacer(strings.NewReplacer(fieldName, paramName))
	}
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

func (loader *Loader[T]) findConfigFilesInEmbedFS(path string, dir embed.FS, ext ConfigExtension) ([]string, error) {
	configsPaths := make([]string, 0)

	err := fs.WalkDir(dir, path, func(path string, info fs.DirEntry, err error) error {
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

func (loader *Loader[T]) findConfigFilesInDir(path string, ext ConfigExtension) ([]string, error) {
	configsPaths := make([]string, 0)

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

func (loader *Loader[T]) toSnakeCase(s string) string {
	var res = make([]rune, 0, len(s))
	var p = '_'
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			res = append(res, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				res = append(res, '_', unicode.ToLower(r))
			} else {
				res = append(res, unicode.ToLower(r))
			}
		} else {
			res = append(res, unicode.ToLower(r))
		}

		p = r
	}
	return string(res)
}
