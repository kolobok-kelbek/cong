package cong

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/loadDirYaml
var embedTestData embed.FS

func Test_Loader_Load(t *testing.T) {
	as := assert.New(t)

	type TestConfig struct {
		ServerName string
		Port       int
		Timeout    int
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.Load("hello", YamlExt, "./testdata/loadYaml")

	as.Nil(err)
	as.NotNil(config)

	as.Equal(config.ServerName, "HelloWorld")
	as.Equal(config.Port, 80)
	as.Equal(config.Timeout, 20)
}

func Test_Loader_Load_withEnvReplace(t *testing.T) {
	as := assert.New(t)

	t.Setenv("HELLO_SERVERNAME", "PAM-PAM")
	t.Setenv("HELLO_PORT", "8080")

	type TestConfig struct {
		ServerName string
		Port       int
		Timeout    int
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.Load("hello", YamlExt, "./testdata/loadYaml")

	as.Nil(err)
	as.NotNil(config)

	as.Equal(config.ServerName, "PAM-PAM")
	as.Equal(config.Port, 8080)
	as.Equal(config.Timeout, 20)
}

func Test_Loader_LoadFromDir(t *testing.T) {
	as := assert.New(t)

	type TestConfig struct {
		App struct {
			Name        string
			Description string
		}
		Db struct {
			Port    int
			Timeout int
		}
		Server struct {
			Name    string
			Port    int
			Timeout int
		}
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromDir("hello", "./testdata/loadDirYaml", YamlExt)

	as.Nil(err)
	as.NotNil(config)

	as.Equal(config.App.Name, "HelloWorld")
	as.Equal(config.App.Description, "This is gorgeous application")
	as.Equal(config.Db.Port, 3036)
	as.Equal(config.Db.Timeout, 10)
	as.Equal(config.Server.Name, "ServerName")
	as.Equal(config.Server.Port, 80)
	as.Equal(config.Server.Timeout, 20)
}

func Test_Loader_LoadFromDir_withEnvReplace(t *testing.T) {
	as := assert.New(t)

	t.Setenv("HELLO_APP_NAME", "APP-NAME")
	t.Setenv("HELLO_SERVER_NAME", "PAM-PAM")
	t.Setenv("HELLO_SERVER_PORT", "8080")

	type TestConfig struct {
		App struct {
			Name        string
			Description string
		}
		Db struct {
			Port    int
			Timeout int
		}
		Server struct {
			Name    string
			Port    int
			Timeout int
		}
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromDir("hello", "./testdata/loadDirYaml", YamlExt)

	as.Nil(err)
	as.NotNil(config)

	as.Equal(config.App.Name, "APP-NAME")
	as.Equal(config.App.Description, "This is gorgeous application")
	as.Equal(config.Db.Port, 3036)
	as.Equal(config.Db.Timeout, 10)
	as.Equal(config.Server.Name, "PAM-PAM")
	as.Equal(config.Server.Port, 8080)
	as.Equal(config.Server.Timeout, 20)
}

func Test_Loader_LoadFromEmbed(t *testing.T) {
	as := assert.New(t)

	type TestConfig struct {
		App struct {
			Name        string
			Description string
		}
		Db struct {
			Port    int
			Timeout int
		}
		Server struct {
			Name    string
			Port    int
			Timeout int
		}
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromEmbedFS("hello", embedTestData, YamlExt)

	as.Nil(err)
	as.NotNil(config)

	as.Equal(config.App.Name, "HelloWorld")
	as.Equal(config.App.Description, "This is gorgeous application")
	as.Equal(config.Db.Port, 3036)
	as.Equal(config.Db.Timeout, 10)
	as.Equal(config.Server.Name, "ServerName")
	as.Equal(config.Server.Port, 80)
	as.Equal(config.Server.Timeout, 20)
}

func Test_Loader_LoadFromEmbed_withEnvReplace(t *testing.T) {
	as := assert.New(t)

	t.Setenv("HELLO_APP_NAME", "APP-NAME")
	t.Setenv("HELLO_SERVER_NAME", "PAM-PAM")
	t.Setenv("HELLO_SERVER_PORT", "8080")

	type TestConfig struct {
		App struct {
			Name        string
			Description string
		}
		Db struct {
			Port    int
			Timeout int
		}
		Server struct {
			Name    string
			Port    int
			Timeout int
		}
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromEmbedFS("hello", embedTestData, YamlExt)

	as.Nil(err)
	as.NotNil(config)

	as.Equal(config.App.Name, "APP-NAME")
	as.Equal(config.App.Description, "This is gorgeous application")
	as.Equal(config.Db.Port, 3036)
	as.Equal(config.Db.Timeout, 10)
	as.Equal(config.Server.Name, "PAM-PAM")
	as.Equal(config.Server.Port, 8080)
	as.Equal(config.Server.Timeout, 20)
}
