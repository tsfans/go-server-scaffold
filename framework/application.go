package framework

import (
	"github.com/tsfans/go/framework/server"
)

var (
	app *Application
)

type Application struct {
	httpServer *server.HttpServer
}

func init() {
	app = &Application{httpServer: server.InitHttpServer()}
}

func Run() {
	app.httpServer.Run()
}
