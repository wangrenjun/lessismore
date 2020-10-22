package log

import (
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

var initconsoleloggeronce sync.Once
var consolelogger zerolog.Logger

func ConsoleLoggerInstance() *zerolog.Logger {
	initconsoleloggeronce.Do(func() {
		consolelogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Caller().Timestamp().Logger()
	})
	return &consolelogger
}
