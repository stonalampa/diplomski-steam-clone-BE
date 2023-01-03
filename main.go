package main

import (
	"fmt"
	"os"

	envs "main/env"

	"github.com/gin-gonic/gin"
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

	// client := database.DatabaseConnector(config)
	// defer client.Disconnect(context.Background())

	if shouldSeed == "seed" {
		// ctx := context.TODO()
		// fmt.Print("OVDE JE@@\n")
		// database.Seeder(env, client, ctx)
	} else {
		fmt.Print(config)
		r := gin.Default()
		r.GET("", func(c *gin.Context) {
			c.String(200, "Welcome to Go and Gin!")
		})
		r.Run(":3030")
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
		// client := database.DatabaseConnector(config)
		// defer cancel()
		// client = dataLayer.InitDataLayer()
		// defer client.Disconnect(context.Background())
		// fmt.Println(db, ctx)
	}

}
