package main

import (
	"encoding/gob"
	"flag"
	"github.com/linchengzhi/go-clean-backend/cmd/initializer"

	"github.com/gin-gonic/gin"
	"github.com/linchengzhi/go-clean-backend/Infra/database"
	"github.com/linchengzhi/go-clean-backend/api/http/middleware"
	"github.com/linchengzhi/go-clean-backend/api/http/routes"
	"github.com/linchengzhi/go-clean-backend/domain/entity"
	"go.uber.org/zap"
)

var configFile = flag.String("f", "../go-clean-backend/config/config_dev.yaml", "the config file")

func main() {
	flag.Parse()

	app, err := initializer.NewApp(*configFile)
	if err != nil {
		panic(err)
	}

	// AutoMigrate
	database.AutoMigrate(app.MysqlDb.DB)

	g := gin.Default()
	gob.Register(entity.User{})
	middleware.SetSessionMiddleware(app.Conf, g)
	routes.SetRoutes(app.UcAll, app.Log, g)
	app.Log.Debug("server is running", zap.Any("config", app.Conf))
	g.Run(":" + app.Conf.HTTP.Port)
}
