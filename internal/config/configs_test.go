package config_test

import (
	"testing"

	"github.com/wangrenjun/lessismore/internal/config"
)

func TestConfigs(t *testing.T) {
	config.Args.ConfigFile = "./configs/lessismore.toml"
	err := config.LoadConfigs()
	if err != nil {
		t.Fatalf("LoadConfigs: %v", err)
	}
	t.Logf("Args: %#v", config.Configs)
}

func TestLoadDotenv(t *testing.T) {
	config.Args.DotEnvFile = "./configs/lessismore.env"
	err := config.LoadDotenv()
	if err != nil {
		t.Fatalf("LoadDotenv: %v", err)
	}
	t.Logf("Args: %#v", config.Configs.DeployEnv)
}
