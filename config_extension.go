package cong

import "github.com/spf13/viper"

type ConfigExtension int

// nolint:revive
const (
	JsonExt ConfigExtension = iota
	TomlExt
	YamlExt
	YmlExt
	PropertiesExt
	PropsExt
	PropExt
	HclExt
	TfvarsExt
	DotenvExt
	EnvExt
	IniExt
)

func (configExtension ConfigExtension) String() string {
	return viper.SupportedExts[configExtension]
}
