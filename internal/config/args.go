package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var Args struct {
	ConfigFile string
	DotEnvFile string
	Version    bool
	Configs    bool
	Routings   bool
	Quiet      bool
}

func ParseArgs() {
	flag.StringVar(&Args.ConfigFile, "config", CONFIG_FILE, "Config File")
	flag.StringVar(&Args.DotEnvFile, "dotenv", DOTENV_FILE, "Dotenv File")
	flag.BoolVar(&Args.Version, "v", false, "Print version")
	flag.BoolVar(&Args.Configs, "c", false, "Dump the configurations")
	flag.BoolVar(&Args.Routings, "r", false, "Dump the routings")
	flag.BoolVar(&Args.Quiet, "q", false, "Quiet Mode")
	flag.Usage = func() {
		color.Set(color.Bold, color.FgGreen)
		defer color.Unset()
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}
