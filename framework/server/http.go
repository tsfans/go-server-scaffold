package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsfans/go/framework/logger"
)

var (
	httpServer *HttpServer
	log        = logger.GetLogger()
)

type HttpServer struct {
	handler *gin.Engine
	srv     *http.Server
}

func InitHttpServer() *HttpServer {
	httpServer = &HttpServer{handler: gin.Default()}
	ip := "0.0.0.0"
	port := "8080"
	address := fmt.Sprintf("%s:%s", ip, port)
	httpServer.srv = &http.Server{
		Addr:           address,
		Handler:        httpServer.handler,
		ReadTimeout:    5 * 60 * 1000,
		WriteTimeout:   5 * 60 * 1000,
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}

func (hs HttpServer) Run() {
	log.Infof("http server is running at %s", hs.srv.Addr)
	if err := hs.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server start error: %s", err)
	}
}
