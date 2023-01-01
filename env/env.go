package env

import (
	"reflect"
)

type Config struct {
	dbConnectionString string
	dbName             string
	dbUsername         string
	dbPassword         string
}

var envConfigs = map[string]Config{
	"localConfig": {
		dbConnectionString: "mongodb://db_user:test123@localhost:27017",
		dbName:             "loyalty-be-db",
		dbUsername:         "db_user",
		dbPassword:         "test123",
	},
	"deployedConfig": {
		dbConnectionString: "mongodb://db_user:test123@localhost:27017",
		dbName:             "loyalty-be-db",
		dbUsername:         "db_user",
		dbPassword:         "test123",
	},
}

func EnvConfig(env string) []string {
	values := reflect.ValueOf(envConfigs[env])
	var array []string
	for i := 0; i < values.NumField(); i++ {
		array = append(array, values.Field(i).String())
	}
	return array
}
