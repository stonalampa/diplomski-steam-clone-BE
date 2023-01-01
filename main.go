package main

import (
	"os"

	models "main/Models"
	envs "main/env"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	env := os.Args[1]
	var config []string

	if env != "local" {
		config = envs.EnvConfig("deployedConfig")
	} else {
		config = envs.EnvConfig("localConfig")
	}

	var dbCredentials = options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    config[1],
		Username:      config[2],
		Password:      config[3],
		PasswordSet:   true,
	}
	_ = mgm.SetDefaultConfig(nil, "loyalty-be-db", options.Client().ApplyURI(config[0]).SetAuth(dbCredentials))

	x := models.User{
		Username:    "SolidStojanXXX",
		Password:    "Test123",
		Name:        "Stole",
		DateOfBirth: "12-03-1999",
	}
	models.CreateUser(&x)
	// models.DeleteUser(&x)
}
