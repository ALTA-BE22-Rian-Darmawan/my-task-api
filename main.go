package main

import (
	"my-task-app/app/config"
	"my-task-app/app/databases"
	"my-task-app/app/migrations"
	"my-task-app/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.RunMigrations(dbMysql)
	// dbPosgres := databases.InitDBPosgres(cfg)

	// create new instance echo
	e := echo.New()

	routes.InitRouter(e, dbMysql)
	e.Use(middleware.Logger())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
