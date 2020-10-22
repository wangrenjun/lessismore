package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Configs struct {
	ListenAddr string
	Log        struct {
		Level      string // https://github.com/rs/zerolog/blob/master/log.go#L134
		Dir        string
		File       string
		MaxBackups int
		MaxSize    int
		MaxAge     int
		Compress   bool
	}
	DeployEnv string
}

func init() {
	Configs.ListenAddr = SERVER_LISTENING_ADDRESS
}

func LoadConfigs() (err error) {
	v := viper.New()
	v.SetConfigFile(Args.ConfigFile)
	if err = v.ReadInConfig(); err != nil {
		return
	}
	err = v.Unmarshal(&Configs)
	return
}

func deployEnv() {
	switch strings.ToLower(os.Getenv("DEPLOY_ENVIRONMENT")) {
	default:
		Configs.DeployEnv = "dev"
	case "dev", "develop", "development":
		Configs.DeployEnv = "dev"
	case "test", "testing":
		Configs.DeployEnv = "test"
	case "product", "production":
		Configs.DeployEnv = "product"
	case "accept", "acceptance":
		Configs.DeployEnv = "accept"
	}
}

func LoadDotenv() error {
	defer deployEnv()
	return godotenv.Load(Args.DotEnvFile)
}
