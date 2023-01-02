package main

import (
	"context"
	"os"

	database "main/database"
	envs "main/env"
)

func main() {
	var shouldSeed string
	var config []string
	env := os.Args[1]
	if len(os.Args) > 2 {
		shouldSeed = os.Args[2]
	}
	if env != "local" {
		config = envs.EnvConfig("deployedConfig")
	} else {
		config = envs.EnvConfig("localConfig")
	}
	client := database.DatabaseConnector(config)
	defer client.Disconnect(context.Background())

	if shouldSeed == "seed" {
		database.Seeder(env, client)
	} else {
		// if env != "local" {
		// 	config = envs.EnvConfig("deployedConfig")
		// } else {
		// 	config = envs.EnvConfig("localConfig")
		// }

		// var db *mongo.Database
		// var ctx context.Context
		// var cancel context.CancelFunc
		// var client *mongo
		//initialize database and context
		client := database.DatabaseConnector(config)
		// defer cancel()
		// client = dataLayer.InitDataLayer()
		defer client.Disconnect(context.Background())
		// fmt.Println(db, ctx)
	}

}
