package seeds

import (
	"context"
	"fmt"
	repo "main/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

func Seeder(db *mongo.Database) {
	for key, element := range Seeds {
		switch v := element.(type) {
		case []repo.User:
			fmt.Printf("Inserting values for type: %s\n", key)
			repo := repo.NewUserRepository(db)
			repo.DropUsers(context.TODO())
			repo.CreateIndices(context.TODO())
			for i := 0; i < len(users); i++ {
				repo.CreateUser(context.TODO(), &users[i])
			}
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
}
