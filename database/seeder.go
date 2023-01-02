package database

import (
	"context"
	models "main/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func Seeder(env string, client *mongo.Client) {
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

	for key, element := range Seeds {
		determineType(key, element, client)
	}
	cancel()
}

func determineType(key string, element any, db *mongo.Database, ctx context.Context) {
	switch key {
	case "Users":
		if users, ok := element.([]models.User); ok {
			for i := 0; i < len(users); i++ {
				models.CreateUser(db, ctx, &users[i])
			}
		}
	case "NeUsers":
		if neusers, ok := element.([]models.User); ok {
			for i := 0; i < len(neusers); i++ {
				models.CreateUser(db, ctx, &neusers[i])
			}
		}
	}
}
