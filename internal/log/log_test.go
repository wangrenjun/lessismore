package log_test

import (
	"testing"

	"github.com/wangrenjun/lessismore/internal/config"

	"github.com/wangrenjun/lessismore/internal/log"
)

func TestLog(t *testing.T) {
	config.Configs.Log.Level = "info"
	config.Configs.Log.File = "test.log"
	config.Configs.Log.Dir = "./"
	config.Configs.Log.MaxBackups = 3
	config.Configs.Log.MaxSize = 1024
	config.Configs.Log.MaxAge = 30
	config.Configs.Log.Compress = true
	log.LoggerInstance().Debug().Str("debug key", "debug val").Msg("debug level")
	log.LoggerInstance().Info().Str("info key", "info val").Msg("info level")
	log.LoggerInstance().Warn().Str("warn key", "warn val").Msg("warn level")
	log.LoggerInstance().Error().Str("error key", "error val").Msg("error level")
	log.LoggerInstance().Fatal().Str("fatal key", "fatal val").Msg("fatal level")
	//log.LoggerInstance().Panic().Str("panic key", "panic val").Msg("panic level")
}
