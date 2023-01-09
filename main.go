package main

import (
	"context"
	"fmt"
	"log"
	"main/repository"
	"main/seeds"
	"main/service"
	"main/utils"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	env     string
	seed    string
	rootCmd = &cobra.Command{
		Use:       "EnvSeed",
		Short:     "Checks env and seed flags",
		Long:      "Checks if the env is local or deployed. Check if the seeds or the server should be run.",
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

	//* Connect to db
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
		repo := repository.NewUserRepository(db)
		userService := service.NewUserService(repo)

		//* Create gin router and set trusted proxy
		router := gin.Default()
		router.SetTrustedProxies([]string{"192.168.0.1"})

		//* Add CORS config to router
		router.Use(cors.Middleware(cors.Config{
			Origins:         "http://files.server.com",
			Methods:         "GET, PUT, DELETE, POST, OPTIONS",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			Credentials:     true,
			ValidateHeaders: true,
		}))

		//* Defined public and private (uses JWT auth) router groups and endpoints
		publicGroup := router.Group("/api")
		privateGroup := router.Group("/api")
		privateGroup.Use(utils.ValidateJwt)
		{
			publicGroup.GET("/users", userService.GetUser)
			privateGroup.POST("/users", userService.CreateUser)
		}

		router.Run(":3030")
	}
}
