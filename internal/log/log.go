package log

import (
	"os"
	"path"
	"sync"
	"time"

	"github.com/wangrenjun/lessismore/internal/config"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var initloggeronce sync.Once
var defaultlogger zerolog.Logger

func LoggerInstance() *zerolog.Logger {
	initloggeronce.Do(func() {
		if err := os.MkdirAll(config.Configs.Log.Dir, 0755); err != nil {
			ConsoleLoggerInstance().Panic().Err(err).Msg("os.MkdirAll error")
		}
		lv, err := zerolog.ParseLevel(config.Configs.Log.Level)
		if err != nil {
			ConsoleLoggerInstance().Panic().Err(err).Msg("zerolog.ParseLevel error")
		}
		zerolog.SetGlobalLevel(lv)
		zerolog.TimeFieldFormat = time.RFC3339
		wr := &lumberjack.Logger{
			Filename:   path.Join(config.Configs.Log.Dir, config.Configs.Log.File),
			MaxBackups: config.Configs.Log.MaxBackups,
			MaxSize:    config.Configs.Log.MaxSize,
			MaxAge:     config.Configs.Log.MaxAge,
			Compress:   config.Configs.Log.Compress,
		}
		defaultlogger = zerolog.New(wr).With().Caller().Timestamp().Logger()
	})
	return &defaultlogger
}
