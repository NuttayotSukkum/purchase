package main

import (
	"context"
	"fmt"
	"github.com/NuttayotSukkum/purchase/configs"
	"github.com/NuttayotSukkum/purchase/configs/db_connector"
	"github.com/NuttayotSukkum/purchase/internal/handlers/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logger "github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	ctx := context.Background()
	cfg := configs.InitConfig(ctx)
	logger.Infof("config :%+v", cfg)
	dbConnector := initDb(cfg)
	logger.Infof("Database is Initializ:%+v", dbConnector)
	rest.ProductRouter(ctx, e, dbConnector)

	execute(ctx, cfg, e)

}

func initDb(cfg *configs.Config) *gorm.DB {
	host := cfg.Secrets.CloudSqlHost
	port := cfg.Secrets.CloudSqlPort
	database := cfg.Secrets.CloudSqlDBName
	username := cfg.Secrets.CloudSqlUser
	password := cfg.Secrets.CloudSqlPass
	return db_connector.NewDb(username, password, host, database, port)
}

func execute(ctx context.Context, cfg *configs.Config, e *echo.Echo) {
	fmt.Sprintf(":%s :%v", cfg.App.Name, cfg.App.Version)
	svcPort := fmt.Sprintf(":%v", cfg.App.Port)
	if err := e.Start(svcPort); err != nil {
		log.Fatal(ctx, "shutting down the server")
	}
}
