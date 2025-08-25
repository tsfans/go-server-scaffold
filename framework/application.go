package framework

import "github.com/gin-gonic/gin"

var (
	app *Application
)

type Application struct {
	WebServer *gin.Engine
}

func init() {
	app = &Application{WebServer: gin.New()}
}

func Run() {
	app.WebServer.Run("0.0.0.0:8080")
}
