package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tsfans/go/framework/config"
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

func NewHttpServer() *HttpServer {
	debug := config.GetBool("debug", true)
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	handler := gin.New()
	handler.ContextWithFallback = true
	handler.Use(gin.Recovery())
	handler.Use(customLogger())
	if trustedProxies := config.GetArray("web.trusted_proxies", []string{}); len(trustedProxies) > 0 {
		if err := handler.SetTrustedProxies(trustedProxies); err != nil {
			log.Panicf("set trusted proxies error: %s", err)
		}
	}

	httpServer = &HttpServer{handler: handler}

	ip := config.GetString("web.listen_host", "0.0.0.0")
	port := config.GetString("web.listen_port", "8080")
	address := fmt.Sprintf("%s:%s", ip, port)
	httpServer.srv = &http.Server{
		Addr:           address,
		Handler:        httpServer.handler,
		ReadTimeout:    time.Duration(config.GetInt("web.timeout.read", 5)) * time.Second,
		WriteTimeout:   time.Duration(config.GetInt("web.timeout.write", 30)) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return httpServer
}

func customLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := float64(time.Since(start)) / float64(time.Millisecond)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		if raw != "" {
			path = path + "?" + raw
		}

		log.Debugf("[http-req-log]%s %s %s %d %0.2fms", clientIP, method, path, statusCode, latency)
	}
}

func (hs HttpServer) Run() {
	go func() {
		log.Infof("http server is running at %s", hs.srv.Addr)
		if err := hs.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server start error: %s", err)
		}
	}()
}

func (hs HttpServer) Shutdown() {
	log.Info("shutting down http server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := hs.srv.Shutdown(ctx); err != nil {
		log.Fatalf("http server forced to shutdown: %s", err)
	}
}
