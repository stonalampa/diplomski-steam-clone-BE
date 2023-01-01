package models

import (
	"fmt"

	"github.com/kamva/mgm/v3"
)

type M map[string]interface{}

// Define our errors:
var internalError = M{"message": "internal error"}
var userNotFound = M{"message": "user not found"}

type User struct {
	mgm.DefaultModel `bson:",inline"`

	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	Name        string `json:"name" bson:"name"`
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth"`
}

// Create handler create new book.
func CreateUser(user *User) int {
	// To get model's collection, just call to mgm.Coll() method.
	err := mgm.Coll(user).Create(user)

	if err != nil {
		// Log your error
		fmt.Print(err)
	}
	return 200
}

func DeleteUser(user *User) int {
	err := mgm.Coll(user).Delete(user)

	if err != nil {
		// Log your error
		fmt.Print(err)
	}
	return 200
}
