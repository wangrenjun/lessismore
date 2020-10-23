package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/gin-contrib/logger"
	"github.com/wangrenjun/lessismore/internal/config"
	"github.com/wangrenjun/lessismore/internal/log"
	"github.com/wangrenjun/lessismore/pkg/utils"

	"github.com/wangrenjun/lessismore/internal/client"
)

func version() {
	if config.Args.Version {
		c := color.New(color.Bold, color.FgYellow)
		c.Printf("GoVersion:                 %s\n", GoVersion)
		c.Printf("SysInfo:                   %s\n", SysInfo)
		c.Printf("LogName:                   %s\n", LogName)
		c.Printf("UserID:                    %s\n", UserID)
		c.Printf("Host:                      %s\n", Host)
		c.Printf("User:                      %s\n", User)
		c.Printf("Email:                     %s\n", Email)
		c.Printf("Repo:                      %s\n", Repo)
		c.Printf("Branch:                    %s\n", Branch)
		c.Printf("LatestTag:                 %s\n", LatestTag)
		c.Printf("LatestCommit:              %s\n", LatestCommit)
		if ts, err := strconv.ParseInt(LatestCommitTimeStamp, 10, 64); err == nil {
			c.Printf("LatestCommitTime:          %s\n", time.Unix(ts, 0).Format(time.RFC3339))
		}
		c.Printf("ModulePath:                %s\n", ModulePath)
		c.Printf("GOOS:                      %s\n", GOOS)
		c.Printf("GOARCH:                    %s\n", GOARCH)
		c.Printf("GOHOSTOS:                  %s\n", GOHOSTOS)
		c.Printf("GOHOSTARCH:                %s\n", GOHOSTARCH)
		c.Printf("SemVer:                    %s\n", SemVer)
		c.Printf("Mode:                      %s\n", Mode)
		if ts, err := strconv.ParseInt(BuildTimeStamp, 10, 64); err == nil {
			c.Printf("BuildTime:                 %s\n", time.Unix(ts, 0).Format(time.RFC3339))
		}
		os.Exit(0)
	}
}

func dump() {
	c := color.New(color.Bold, color.FgYellow)
	c.Printf("PID:                       %d\n", os.Getpid())
	c.Printf("Args:                      %s\n", utils.PrettyJson(config.Args))
	c.Printf("Configs:                   %s\n", utils.PrettyJson(config.Configs))
	c.Printf("DEPLOY_ENVIRONMENT:        %s\n", os.Getenv("DEPLOY_ENVIRONMENT"))
	c.Printf("DB_DRIVER:                 %s\n", os.Getenv("DB_DRIVER"))
	c.Printf("DB_USER:                   %s\n", os.Getenv("DB_USER"))
	c.Printf("DB_PORT:                   %s\n", os.Getenv("DB_PORT"))
	c.Printf("DB_HOST:                   %s\n", os.Getenv("DB_HOST"))
	c.Printf("DB_NAME:                   %s\n", os.Getenv("DB_NAME"))
	c.Printf("REDIS_HOST:                %s\n", os.Getenv("REDIS_HOST"))
	c.Printf("REDIS_PORT:                %s\n", os.Getenv("REDIS_PORT"))
	client.PathRouterInstance().Range(func(pattern string, handler client.WsHandlerFunc) {
		color.New(color.Bold, color.FgYellow).Printf("Registered handler:        Path: %s, Handler: %s\n", pattern, utils.GetFunctionName(handler))
	})
}

func Run() {
	config.ParseArgs()
	version()
	if err := config.LoadConfigs(); err != nil {
		log.ConsoleLoggerInstance().Fatal().Err(err).Msg("LoadConfigs error")
	}
	if err := config.LoadDotenv(); err != nil {
		log.ConsoleLoggerInstance().Fatal().Err(err).Msg("LoadDotenv error")
	}
	dump()
	router := client.HttpRouterInstance()
	router.Use(logger.SetLogger(logger.Config{
		Logger: log.LoggerInstance(),
	}))
	server := &http.Server{
		Addr:    config.Configs.ListenAddr,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.ConsoleLoggerInstance().Panic().Err(err).Msg("ListenAndServe failed")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		sig := <-quit
		log.LoggerInstance().Info().Str("Signal", sig.String()).Msg("Signal received")
		if sig == syscall.SIGINT {
			utils.Tempfile("./", "STACKTRACE-"+time.Now().Format("20060102150405"), utils.StackTrace(true))
			continue
		}
		break
	}
	log.LoggerInstance().Info().Msg("Server shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.LoggerInstance().Fatal().Err(err).Msg("Server shutdown by forced")
	}
}
