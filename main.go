package main

import (
	"mia/my_task_app/apps/config"
	"mia/my_task_app/apps/database"
	"mia/my_task_app/apps/router"

	"github.com/labstack/echo/v4"
)

func main() {
	// logging := helpers.NewLogger()
	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)

	//call migration
	database.InitialMigration(dbMysql)

	//create a new echo instance
	e := echo.New()
	router.InitRouter(dbMysql, e)

	//start server and port
	e.Logger.Fatal(e.Start(":8083"))
}
