package seeds

import (
	repo "main/repository"
	"main/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var libraryRecords = []repo.LibraryRecord{
	{
		ID:          utils.GenerateId("63bd6bd756fa7318db852016"),
		UserId:      utils.GenerateId("63bd69c3b2058767688d3d94"),
		GameIds:     []primitive.ObjectID{utils.GenerateId("63bed8174d991d50de422d6b"), utils.GenerateId("63bed81ec1bd88d167f66a8f")},
		WishlistIds: []primitive.ObjectID{utils.GenerateId("63bed80f04bedac5324c421a")},
	},
	{
		ID:          utils.GenerateId("63bef748af78078b25210645"),
		UserId:      utils.GenerateId("63bd6bd756fa7318db852016"),
		GameIds:     []primitive.ObjectID{},
		WishlistIds: []primitive.ObjectID{},
	},
	{
		ID:          utils.GenerateId("63bef62e5edc158b2dae28cb"),
		UserId:      utils.GenerateId("63bd6d1bd0564a6e1c2de5b4"),
		GameIds:     []primitive.ObjectID{},
		WishlistIds: []primitive.ObjectID{utils.GenerateId("63bed8230a995357e03daa9f"), utils.GenerateId("63bed828e88e68df55def4e4")},
	},
	{
		ID:          utils.GenerateId("63bef6329be7376a67d8a375"),
		UserId:      utils.GenerateId("63bd6d1fab372bf383a2dbc7"),
		GameIds:     []primitive.ObjectID{utils.GenerateId("63bed8377f1a60b806b81cb8"), utils.GenerateId("63bed843597757fc6bb4953e"), utils.GenerateId("63bed848240a8702cbdf09d4")},
		WishlistIds: []primitive.ObjectID{},
	},
	{
		ID:          utils.GenerateId("63bef687007c41c2eaff43d6"),
		UserId:      utils.GenerateId("63bd6d235a1cbb4302d1f09a"),
		GameIds:     []primitive.ObjectID{utils.GenerateId("63bed843597757fc6bb4953e"), utils.GenerateId("63bed83dc0378d4abb338e72")},
		WishlistIds: []primitive.ObjectID{utils.GenerateId("63bed83259e1a173be264188"), utils.GenerateId("63bed8377f1a60b806b81cb8")},
	},
	{
		ID:          utils.GenerateId("63bef62af5f5411fb796663c"),
		UserId:      utils.GenerateId("63bd6d29bfdab9586c5bf162"),
		GameIds:     []primitive.ObjectID{utils.GenerateId("63bed83259e1a173be264188"), utils.GenerateId("63bed83dc0378d4abb338e72")},
		WishlistIds: []primitive.ObjectID{utils.GenerateId("63bed843597757fc6bb4953e"), utils.GenerateId("63bed8377f1a60b806b81cb8")},
	},
	{
		ID:          utils.GenerateId("63bef6276b2c9b72bd79c1c3"),
		UserId:      utils.GenerateId("63bd6d40fb221e7af31f30dd"),
		GameIds:     []primitive.ObjectID{utils.GenerateId("63bed8230a995357e03daa9f"), utils.GenerateId("63bed80f04bedac5324c421a")},
		WishlistIds: []primitive.ObjectID{utils.GenerateId("63bed848240a8702cbdf09d4"), utils.GenerateId("63bed83259e1a173be264188")},
	},
	{
		ID:          utils.GenerateId("63bef623b376dd5533606212"),
		UserId:      utils.GenerateId("63bd6ce339dd5484c8d7b7a9"),
		GameIds:     []primitive.ObjectID{utils.GenerateId("63bed81ec1bd88d167f66a8f"), utils.GenerateId("63bed8174d991d50de422d6b")},
		WishlistIds: []primitive.ObjectID{utils.GenerateId("63bed8230a995357e03daa9f"), utils.GenerateId("63bed848240a8702cbdf09d4")},
	},
	{
		ID:          utils.GenerateId("63bef61f28995bcd976bd3bd"),
		UserId:      utils.GenerateId("63bd6d3408644f39445e6ee0"),
		GameIds:     []primitive.ObjectID{utils.GenerateId("63bed80f04bedac5324c421a"), utils.GenerateId("63bed848240a8702cbdf09d4")},
		WishlistIds: []primitive.ObjectID{utils.GenerateId("63bed83259e1a173be264188"), utils.GenerateId("63bed8174d991d50de422d6b")},
	},
}
