package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tsfans/go/framework/logger"
)

var (
	httpServer *HttpServer
	log        = logger.GetLogger()
)

type HttpServer struct {
	*gin.Engine
	listenAddr string
}

func InitHttpServer() *HttpServer {
	httpServer = &HttpServer{Engine: gin.Default(), listenAddr: ":8080"}
	return httpServer
}

func (hs HttpServer) Run() {
	log.Infof("http server is running at %s", hs.listenAddr)
	hs.Engine.Run(hs.listenAddr)
}
