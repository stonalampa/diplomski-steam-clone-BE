package main

import (
	"context"
	"fmt"
	"log"
	"main/seeds"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	env     string
	seed    string
	rootCmd = &cobra.Command{
		Use:       "seed",
		Short:     "If TRUE it will run seeding.",
		Long:      "If TRUE the seeding will be done, if FALSE the program will run normally.",
		ValidArgs: []string{"true", "false", "local", "deployment"},
		Args:      matchAll(cobra.MinimumNArgs(2), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "true" || args[0] == "false" {
				env = args[1]
				seed = args[0]
			} else {
				env = args[0]
				seed = args[1]
			}
		},
	}
)

func matchAll(checks ...cobra.PositionalArgs) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		for _, check := range checks {
			if err := check(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//* Set configuration
	configLocation := "./env/" + env + ".yaml"
	viper.SetConfigFile(configLocation)
	viper.Set("env", env)
	seedBoolean, booleanErr := strconv.ParseBool(seed)
	if booleanErr != nil {
		log.Fatal(booleanErr)
	}
	viper.Set("seed", seedBoolean)
	configError := viper.ReadInConfig()
	if configError != nil {
		panic(fmt.Errorf("fatal error config file: %w", configError))
	}

	//* Connect and create db
	var dbCredentials = options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    viper.GetString("name"),
		Username:      viper.GetString("username"),
		Password:      viper.GetString("password"),
		PasswordSet:   true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("connectionString")).SetAuth((dbCredentials)))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(viper.GetString("name"))

	if viper.GetBool("seed") == true {
		seeds.Seeder(db)
	} else {
		//RUN SERVER
	}

	// * Connect to server
	// client := database.DatabaseConnector(config)
	// defer client.Disconnect(context.Background())

	// if shouldSeed == "seed" {
	// 	// ctx := context.TODO()
	// 	// fmt.Print("OVDE JE@@\n")
	// 	// database.Seeder(env, client, ctx)
	// } else {
	// 	fmt.Print(config)
	// 	r := gin.Default()
	// 	r.GET("", func(c *gin.Context) {
	// 		c.String(200, "Welcome to Go and Gin!")
	// 	})
	// 	r.Run(":3030")
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
	// }

}
