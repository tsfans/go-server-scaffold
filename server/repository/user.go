package repository

import (
	"github.com/tsfans/go/framework/config"
	"github.com/tsfans/go/framework/database"
	"github.com/tsfans/go/server/model/po"
)

func init() {
	migrate := config.GetBool("postgre.auto_migrate.user", false)
	if !migrate {
		return
	}
	database.DB.AutoMigrate(&po.User{})
}
