package seeds

import (
	"context"
	"fmt"
	repo "main/repository"
	"main/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var seeds = map[string]any{
	"Users":          users,
	"Games":          games,
	"LibraryRecords": libraryRecords,
	"Reviews":        reviews,
}

func Seeder(db *mongo.Database) {
	for key, element := range seeds {
		switch v := element.(type) {
		case []repo.User:
			printMessage(key)
			repo := repo.NewUserRepository(db)
			repo.DropUsers(context.TODO())
			repo.CreateIndices(context.TODO())
			for i := 0; i < len(users); i++ {
				repo.CreateUser(context.TODO(), &users[i])
			}
		case []repo.Game:
			printMessage(key)
		case []repo.LibraryRecord:
			printMessage(key)
		case []repo.Review:
			printMessage(key)
		default:
			fmt.Printf("Unknown type %T!\n", v)
		}
	}
}

func printMessage(key string) {
	fmt.Printf("Inserting values for type: %s\n", key)
}

func generateId(hex string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(hex)
	return id
}

func generatePassword(pass string) string {
	hashedPass, _ := utils.HashPassword(pass)
	return hashedPass
}
