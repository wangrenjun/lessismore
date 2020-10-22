package log_test

import (
	"testing"

	"github.com/wangrenjun/lessismore/internal/log"
)

func TestConsole(t *testing.T) {
	log.ConsoleLoggerInstance().Debug().Str("debug key", "debug val").Msg("debug level")
	log.ConsoleLoggerInstance().Error().Str("error key", "error val").Msg("error level")
	//log.ConsoleLoggerInstance().Panic().Str("panic key", "panic val").Msg("panic level")
}
