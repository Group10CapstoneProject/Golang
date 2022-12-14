package main

import (
	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/database"
	"github.com/Group10CapstoneProject/Golang/route"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitConfig()

	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	if err := database.MigrateDB(db); err != nil {
		panic(err)
	}

	e := echo.New()

	route.InitRoutes(e, db)
	e.Logger.Fatal(e.Start(":" + config.Env.API_PORT))
}
