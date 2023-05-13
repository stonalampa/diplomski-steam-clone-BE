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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		ValidArgs: []string{"true", "false", "local", "deployment", ""},
		Args:      matchAll(cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				env = "deployment"
				seed = "false"
			} else if args[0] == "true" || args[0] == "false" {
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

	// * Set configuration
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

	// * Connect to db
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

	if viper.GetBool("seed") {
		seeds.Seeder(db)
	} else {
		userRepo := repository.NewUsersRepository(db)
		authService := service.NewAuthService(userRepo)
		userService := service.NewUsersService(userRepo)

		gamesRepo := repository.NewGamesRepository(db)
		gamesService := service.NewGamesService(gamesRepo)

		libraryRepo := repository.NewLibraryRepository(db)
		libraryService := service.NewLibraryService(libraryRepo)

		reviewRepo := repository.NewReviewsRepository(db)
		reviewService := service.NewReviewsService(reviewRepo)

		// * Create gin router and set trusted proxy
		router := gin.Default()
		router.SetTrustedProxies([]string{"192.168.0.1"})

		// * Add CORS config to router
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"https://localhost:8080", "http://localhost:8080"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Content-Length", "Content-Type", "authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		// * Defined public and private (uses JWT auth) router groups and endpoints
		publicGroup := router.Group("/api")
		privateGroup := router.Group("/api")
		privateGroup.Use(utils.ValidateJwt)
		{
			// * Login
			publicGroup.POST("/adminLogin", authService.AdminLogin)
			publicGroup.POST("/login", authService.Login)

			// * Users
			privateGroup.GET("/users/:id", userService.GetUser)
			privateGroup.GET("/users", userService.GetUsers)
			privateGroup.POST("/users", userService.CreateUser)
			privateGroup.PUT("/users", userService.UpdateUser)
			privateGroup.DELETE("/users", userService.DeleteUser)

			// * Games
			publicGroup.GET("/games", gamesService.GetAllGames)
			publicGroup.GET("/games/:id", gamesService.GetGame)

			privateGroup.POST("/games", gamesService.CreateGame)
			privateGroup.PUT("/games", gamesService.UpdateGame)
			privateGroup.DELETE("/games", gamesService.DeleteGame)

			// * Library
			privateGroup.GET("/library/:id", libraryService.GetLibraryRecord)
			privateGroup.POST("/library", libraryService.CreateLibraryRecord)
			privateGroup.PUT("/library", libraryService.UpdateLibraryRecord)
			privateGroup.DELETE("/library", libraryService.DeleteLibraryRecord)

			// * Reviews
			privateGroup.GET("/reviews/:id", reviewService.GetReviewRecord)
			privateGroup.POST("/reviews", reviewService.CreateReviewRecord)
			privateGroup.PUT("/reviews", reviewService.UpdateReviewRecord)
			privateGroup.DELETE("/reviews", reviewService.DeleteReviewRecord)
		}

		router.Run(":3030")
	}
}
