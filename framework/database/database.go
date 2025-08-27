package database

import (
	"github.com/tsfans/go/framework/config"
	"github.com/tsfans/go/framework/logger"
)

var (
	log = logger.Get()
)

func init() {
	log.Debug("initializing database ...")
	if config.Exists("postgre") {
		initPostgreDB()
	}
}
