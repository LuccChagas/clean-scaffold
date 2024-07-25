package main

import (
	"github.com/LuccChagas/clean-scaffold/config"
	"log"
)

// TODO: Need implement zap logger as global

func main() {

	var err error
	_, err = config.SetupEnv()
	if err != nil {
		panic(err)
	}

	sqlDB, err := config.ConnDB()
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	app := config.NewApp(sqlDB)
	//logger.Sync()
	app.Server.Serve()

	log.Println("App Up and Running")

}
