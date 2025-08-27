package framework

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/tsfans/go/framework/config"
	"github.com/tsfans/go/framework/logger"
	"github.com/tsfans/go/framework/server"
)

var (
	app *Application
	log = logger.Get()
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

func LoadRoute(loader func(c *gin.RouterGroup)) {
	engine := app.httpServer.GetRouterEngine()
	loader(&engine.RouterGroup)
}

func LoadServerRoute(loader func(c *gin.RouterGroup), version ...string) {
	engine := app.httpServer.GetRouterEngine()
	serverName := config.GetString("server.name", "go-server-scaffold")
	prefix := config.GetString("web.route_prefix", serverName)
	loader(engine.Group(prefix))
}
