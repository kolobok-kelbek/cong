package cong

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/loadDirYaml
var embedTestData embed.FS

//go:embed testdata
var embedFullTestData embed.FS

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

	as.Equal(config, &TestConfig{
		ServerName: "HelloWorld",
		Port:       80,
		Timeout:    20,
	})
}

func Test_Loader_Load_withEnvReplace(t *testing.T) {
	as := assert.New(t)

	t.Setenv("HELLO_SERVER_NAME", "PAM-PAM")
	t.Setenv("HELLO_PORT", "8080")

	type TestConfig struct {
		ServerName string
		Port       int
		Timeout    int
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.Load("hello", YamlExt, "./testdata/loadYaml")

	as.Nil(err)
	as.Equal(config, &TestConfig{
		ServerName: "PAM-PAM",
		Port:       8080,
		Timeout:    20,
	})
}

func Test_Loader_LoadFromDir(t *testing.T) {
	as := assert.New(t)

	type App struct {
		Name        string
		Description string
	}
	type Db struct {
		Port    int
		Timeout int
	}
	type Server struct {
		Name    string
		Port    int
		Timeout int
	}
	type TestConfig struct {
		App    App
		Db     Db
		Server Server
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromDir("hello", "./testdata/loadDirYaml", YamlExt)

	as.Nil(err)
	as.Equal(config, &TestConfig{
		App: App{
			Name:        "HelloWorld",
			Description: "This is gorgeous application",
		},
		Db: Db{
			Port:    3036,
			Timeout: 10,
		},
		Server: Server{
			Name:    "ServerName",
			Port:    80,
			Timeout: 20,
		},
	})
}

func Test_Loader_LoadFromDir_withEnvReplace(t *testing.T) {
	as := assert.New(t)

	t.Setenv("HELLO_APP_NAME", "APP-NAME")
	t.Setenv("HELLO_SERVER_NAME", "PAM-PAM")
	t.Setenv("HELLO_SERVER_PORT", "8080")

	type App struct {
		Name        string
		Description string
	}
	type Db struct {
		Port    int
		Timeout int
	}
	type Server struct {
		Name    string
		Port    int
		Timeout int
	}
	type TestConfig struct {
		App    App
		Db     Db
		Server Server
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromDir("hello", "./testdata/loadDirYaml", YamlExt)

	as.Nil(err)
	as.Equal(config, &TestConfig{
		App: App{
			Name:        "APP-NAME",
			Description: "This is gorgeous application",
		},
		Db: Db{
			Port:    3036,
			Timeout: 10,
		},
		Server: Server{
			Name:    "PAM-PAM",
			Port:    8080,
			Timeout: 20,
		},
	})
}

func Test_Loader_LoadFromEmbed(t *testing.T) {
	as := assert.New(t)

	type App struct {
		Name        string
		Description string
	}
	type Db struct {
		Port    int
		Timeout int
	}
	type Server struct {
		Name    string
		Port    int
		Timeout int
	}
	type TestConfig struct {
		App    App
		Db     Db
		Server Server
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromEmbedFS("hello", embedTestData, YamlExt)

	as.Nil(err)
	as.Equal(config, &TestConfig{
		App: App{
			Name:        "HelloWorld",
			Description: "This is gorgeous application",
		},
		Db: Db{
			Port:    3036,
			Timeout: 10,
		},
		Server: Server{
			Name:    "ServerName",
			Port:    80,
			Timeout: 20,
		},
	})
}

func Test_Loader_LoadFromEmbed_withEnvReplace(t *testing.T) {
	as := assert.New(t)

	t.Setenv("HELLO_APP_NAME", "APP-NAME")
	t.Setenv("HELLO_SERVER_NAME", "PAM-PAM")
	t.Setenv("HELLO_SERVER_PORT", "8080")

	type App struct {
		Name        string
		Description string
	}
	type Db struct {
		Port    int
		Timeout int
	}
	type Server struct {
		Name    string
		Port    int
		Timeout int
	}
	type TestConfig struct {
		App    App
		Db     Db
		Server Server
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromEmbedFS("hello", embedTestData, YamlExt)

	as.Nil(err)
	as.Equal(config, &TestConfig{
		App: App{
			Name:        "APP-NAME",
			Description: "This is gorgeous application",
		},
		Db: Db{
			Port:    3036,
			Timeout: 10,
		},
		Server: Server{
			Name:    "PAM-PAM",
			Port:    8080,
			Timeout: 20,
		},
	})
}

func Test_Loader_LoadFromEmbedByPath(t *testing.T) {
	as := assert.New(t)

	type App struct {
		Name        string
		Description string
	}
	type Db struct {
		Port    int
		Timeout int
	}
	type Server struct {
		Name    string
		Port    int
		Timeout int
	}
	type TestConfig struct {
		App    App
		Db     Db
		Server Server
	}

	loader := NewLoader[TestConfig]()

	config, err := loader.LoadFromEmbedFSByPath("hello", embedFullTestData, "testdata/loadDirYaml", YamlExt)

	as.Nil(err)
	as.Equal(config, &TestConfig{
		App: App{
			Name:        "HelloWorld",
			Description: "This is gorgeous application",
		},
		Db: Db{
			Port:    3036,
			Timeout: 10,
		},
		Server: Server{
			Name:    "ServerName",
			Port:    80,
			Timeout: 20,
		},
	})
}

func TestLoader_Load_HugeFile(t *testing.T) {
	as := assert.New(t)

	type Level5Config struct {
		Key        string   `mapstructure:"key"`
		Pi         float64  `mapstructure:"pi"`
		Enabled    bool     `mapstructure:"enabled"`
		List       []string `mapstructure:"list"`
		AnotherKey string   `mapstructure:"anotherKey"`
	}

	type Level4Config struct {
		Level5 Level5Config `mapstructure:"level5"`
	}

	type Level3Config struct {
		Level4 Level4Config `mapstructure:"level4"`
	}

	type Level2Config struct {
		Level3 Level3Config `mapstructure:"level3"`
	}

	type Level1Config struct {
		Level2 Level2Config `mapstructure:"level2"`
	}

	type NestedConfig struct {
		Level1 Level1Config `mapstructure:"level1"`
	}

	type OAuthConfig struct {
		ClientID     string `mapstructure:"clientId"`
		ClientSecret string `mapstructure:"clientSecret"`
	}

	type ThirdPartyConfig struct {
		Google   OAuthConfig `mapstructure:"google"`
		Facebook OAuthConfig `mapstructure:"facebook"`
	}

	type LogFile struct {
		Path       string `mapstructure:"path"`
		MaxSize    string `mapstructure:"maxsize"`
		MaxBackups int    `mapstructure:"maxbackups"`
		MaxAge     string `mapstructure:"maxage"`
	}

	type LogConfig struct {
		Level  string   `mapstructure:"level"`
		Format string   `mapstructure:"format"`
		Output []string `mapstructure:"output"`
		File   LogFile  `mapstructure:"file"`
	}

	type SSLConfig struct {
		Enabled         bool   `mapstructure:"enabled"`
		CertificatePath string `mapstructure:"certificatePath"`
		PrivateKeyPath  string `mapstructure:"privateKeyPath"`
	}

	type CorsConfig struct {
		Enabled        bool     `mapstructure:"enabled"`
		AllowedOrigins []string `mapstructure:"allowedOrigins"`
		MaxAge         string   `mapstructure:"maxAge"`
	}

	type ServerConfig struct {
		Host string     `mapstructure:"host"`
		Port int        `mapstructure:"port"`
		Cors CorsConfig `mapstructure:"cors"`
		SSL  SSLConfig  `mapstructure:"ssl"`
	}

	type DatabaseConfig struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		Timeout  string `mapstructure:"timeout"`
	}

	type Configuration struct {
		Database      DatabaseConfig   `mapstructure:"database"`
		Server        ServerConfig     `mapstructure:"server"`
		Logs          LogConfig        `mapstructure:"logs"`
		ThirdParty    ThirdPartyConfig `mapstructure:"thirdParty"`
		NestedExample NestedConfig     `mapstructure:"nestedExample"`
	}

	type AppMetadata struct {
		Author string `mapstructure:"author"`
		Email  string `mapstructure:"email"`
		Year   int    `mapstructure:"year"`
	}

	type AppConfig struct {
		Name        string        `mapstructure:"name"`
		Version     string        `mapstructure:"version"`
		Description string        `mapstructure:"description"`
		Metadata    AppMetadata   `mapstructure:"metadata"`
		Features    []string      `mapstructure:"features"`
		Config      Configuration `mapstructure:"config"`
	}

	type App struct {
		App AppConfig `mapstructure:"app"`
	}

	loader := NewLoader[App]()

	config, err := loader.Load("huge", YamlExt, "./testdata/loadYaml")

	as.Nil(err)
	as.Equal(config, &App{
		App: AppConfig{
			Name:        "My Super Application",
			Version:     "v1.2.3",
			Description: "This is a very long description of the application. It spans multiple lines and is intended to give an overview of what the application is and does.",
			Metadata: AppMetadata{
				Author: "Jane Doe",
				Email:  "jane.doe@example.com",
				Year:   2023,
			},
			Features: []string{"authentication", "logging", "analytics"},
			Config: Configuration{
				Database: DatabaseConfig{
					Username: "dbuser",
					Password: "securepassword",
					Host:     "localhost",
					Port:     5432,
					Name:     "mydatabase",
					Timeout:  "10s",
				},
				Server: ServerConfig{
					Host: "0.0.0.0",
					Port: 8080,
					Cors: CorsConfig{
						Enabled:        true,
						AllowedOrigins: []string{"https://example1.com", "https://example2.com"},
						MaxAge:         "300s",
					},
					SSL: SSLConfig{
						Enabled:         false,
						CertificatePath: "/etc/ssl/certs/cert.pem",
						PrivateKeyPath:  "/etc/ssl/private/key.pem",
					},
				},
				Logs: LogConfig{
					Level:  "info",
					Format: "json",
					Output: []string{"stdout", "file"},
					File: LogFile{
						Path:       "/var/log/myapp.log",
						MaxSize:    "100MB",
						MaxBackups: 5,
						MaxAge:     "30d",
					},
				},
				ThirdParty: ThirdPartyConfig{
					Google: OAuthConfig{
						ClientID:     "google-client-id",
						ClientSecret: "google-client-secret",
					},
					Facebook: OAuthConfig{
						ClientID:     "facebook-app-id",
						ClientSecret: "facebook-app-secret",
					},
				},
				NestedExample: NestedConfig{
					Level1: Level1Config{
						Level2: Level2Config{
							Level3: Level3Config{
								Level4: Level4Config{
									Level5: Level5Config{
										Key:        "deeply nested value",
										Pi:         3.14159265359,
										Enabled:    true,
										List:       []string{"item1", "item2", "item3"},
										AnotherKey: "This is another long text that is associated with 'anotherKey' in this deeply nested structure.\n",
									},
								},
							},
						},
					},
				},
			},
		},
	})
}
