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
			repo := repo.NewUsersRepository(db)
			repo.DropUsers(context.TODO())
			repo.CreateIndices(context.TODO())
			for i := 0; i < len(users); i++ {
				repo.CreateUser(context.TODO(), &users[i])
			}
		case []repo.Game:
			printMessage(key)
			repo := repo.NewGamesRepository(db)
			repo.DropGames(context.TODO())
			repo.CreateIndices(context.TODO())
			for i := 0; i < len(games); i++ {
				repo.CreateGame(context.TODO(), &games[i])
			}
		case []repo.LibraryRecord:
			printMessage(key)
			repo := repo.NewLibraryRepository(db)
			repo.DropLibraryRecords(context.TODO())
			repo.CreateIndices(context.TODO())
			for i := 0; i < len(libraryRecords); i++ {
				repo.CreateLibraryRecord(context.TODO(), &libraryRecords[i])
			}
		case []repo.Review:
			printMessage(key)
			repo := repo.NewReviewsRepository(db)
			repo.DropReviews(context.TODO())
			repo.CreateIndices(context.TODO())
			for i := 0; i < len(reviews); i++ {
				repo.CreateReview(context.TODO(), &reviews[i])
			}
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
