package database

import (
	"context"
	"fmt"
	models "main/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func Seeder(env string, client *mongo.Client, ctx context.Context) {
	// var config []string
	// if env != "local" {
	// 	config = envs.EnvConfig("deployedConfig")
	// } else {
	// 	config = envs.EnvConfig("localConfig")
	// }
	// var db *mongo.Database
	// var ctx context.Context
	// var cancel context.CancelFunc
	//initialize database and context
	// db, ctx, cancel = DatabaseConnector(config)
	// ctx := context.TODO()
	for key, element := range Seeds {
		determineType(key, element, ctx)
	}
	// cancel()
}

func determineType(key string, element any, ctx context.Context) {
	fmt.Print("OVE U DETERMEINE\n", ctx)
	switch key {
	case "Users":
		if users, ok := element.([]models.User); ok {
			var userRepo models.UserRepository
			for i := 0; i < len(users); i++ {
				fmt.Print("OVDE U FOR")
				t, err := userRepo.CreateUser(ctx, users[i])
				if err != nil {
					fmt.Print(err)
				}
				fmt.Println(t)
			}
		}
		// case "NeUsers":
		// 	if neusers, ok := element.([]models.User); ok {
		// 		var userRepo models.UserRepository
		// 		for i := 0; i < len(neusers); i++ {
		// 			userRepo.CreateNeUser(ctx, neusers[i])
		// 		}
		// 	}
	}
}
