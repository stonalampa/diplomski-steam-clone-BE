package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
* config is array of strings created from Config type located in env.go file
* config[0] - dbConnectionString
* config[1] - dbName
* config[2] - dbUsername
* config[3] - dbPassword
 */
// func DatabaseConnector(config []string) (*mongo.Database, context.Context, context.CancelFunc) {
func DatabaseConnector(config []string) *mongo.Database {
	var dbCredentials = options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    config[1],
		Username:      config[2],
		Password:      config[3],
		PasswordSet:   true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config[0]).SetAuth((dbCredentials)))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
	return client.Database(config[1])
}
