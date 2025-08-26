package framework

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/tsfans/go/framework/config"
	"github.com/tsfans/go/framework/logger"
	"github.com/tsfans/go/framework/server"
)

var (
	app *Application
	log = logger.GetLogger()
)

type Application struct {
	httpServer *server.HttpServer
}

func init() {
	logLevel := config.GetString("log.level", "debug")
	logFile := filepath.Join(config.GetString("log.path", ""), config.GetString("log.file", "app.log"))
	logger.InitLogger(logLevel, logFile)
	app = &Application{httpServer: server.NewHttpServer()}
}

func Run() {
	app.httpServer.Run()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for sig := range exit {
		log.Infof("received exit signal: %s", sig.String())
		app.httpServer.Shutdown()
	}
}
