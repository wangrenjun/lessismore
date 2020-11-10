package client

import (
	"expvar"
	"runtime"
	"time"

	ginexpvar "github.com/gin-contrib/expvar"
)

type uptime struct {
	start time.Time
}

func (u uptime) String() string {
	return time.Since(u.start).String()
}

func init() {
	expvar.Publish("uptime", &uptime{start: time.Now()})
	expvar.NewString("GoVersion").Set(runtime.Version())
	expvar.NewString("GOOS").Set(runtime.GOOS)
	expvar.NewString("GOARCH").Set(runtime.GOARCH)
	expvar.NewInt("GOMAXPROCS").Set(int64(runtime.GOMAXPROCS(-1)))
	expvar.NewInt("NumCPU").Set(int64(runtime.NumCPU()))
	expvar.NewInt("NumGoroutine").Set(int64(runtime.NumGoroutine()))
	expvar.NewInt("NumCgoCall").Set(runtime.NumCgoCall())

	HttpRouterInstance().GET("/debug/vars", ginexpvar.Handler())
}
