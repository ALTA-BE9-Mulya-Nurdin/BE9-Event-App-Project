package main

import (
	"be9/event/config"
	"be9/event/factory"
	_middlewares "be9/event/middlewares"
	"be9/event/migration"
	"be9/event/routes"
)

func main() {
	dbConn := config.InitDB()
	migration.InitMigrate(dbConn)
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	_middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
