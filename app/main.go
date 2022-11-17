package main

import (
	"fmt"
	"time"

	"github.com/Group10CapstoneProject/Golang/config"
	"github.com/Group10CapstoneProject/Golang/database"
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

	fmt.Println("success initial")
	fmt.Println(time.Now())
}
