package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LibraryRepository interface {
	CreateLibraryRecord(ctx context.Context, libraryRecord *LibraryRecord) (*mongo.InsertOneResult, error)
	GetLibraryRecord(ctx context.Context, userId primitive.ObjectID) (LibraryRecord, error)
	UpdateLibraryRecord(ctx context.Context, libraryRecord *LibraryRecord, remove bool) (*mongo.UpdateResult, error)
	DeleteLibraryRecord(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
	DropLibraryRecords(ctx context.Context)
	CreateIndices(ctx context.Context)
}

type libraryRepository struct {
	db *mongo.Database
}

func NewLibraryRepository(db *mongo.Database) LibraryRepository {
	return &libraryRepository{db: db}
}

type LibraryRecord struct {
	ID          primitive.ObjectID   `bson:"_id"`
	UserId      primitive.ObjectID   `json:"userId" bson:"userId"`
	GameIds     []primitive.ObjectID `json:"gameIds" bson:"gameIds"`
	WishlistIds []primitive.ObjectID `json:"wishlistIds" bson:"wishlistIds"`
}

func (repo *libraryRepository) CreateLibraryRecord(ctx context.Context, libraryRecord *LibraryRecord) (*mongo.InsertOneResult, error) {
	result, err := repo.db.Collection("library").InsertOne(ctx, libraryRecord)
	return result, err
}

func (repo *libraryRepository) GetLibraryRecord(ctx context.Context, userId primitive.ObjectID) (LibraryRecord, error) {
	var libRecord LibraryRecord
	filter := bson.D{primitive.E{Key: "userId", Value: userId}}
	err := repo.db.Collection("library").FindOne(ctx, filter).Decode(&libRecord)
	if err != nil {
		return LibraryRecord{}, err
	}

	return libRecord, nil
}

func (repo *libraryRepository) UpdateLibraryRecord(ctx context.Context, libraryRecord *LibraryRecord, remove bool) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": libraryRecord.ID, "userId": libraryRecord.UserId}
	update := bson.M{}

	if len(libraryRecord.GameIds) > 0 {
		if !remove {
			update["$pull"] = bson.M{"wishlistIds": bson.M{"$in": libraryRecord.GameIds}}
		}

		if remove {
			update["$pull"] = bson.M{"gameIds": bson.M{"$in": libraryRecord.GameIds}}
		} else {
			update["$addToSet"] = bson.M{"gameIds": bson.M{"$each": libraryRecord.GameIds}}
		}
	}

	if len(libraryRecord.WishlistIds) > 0 {
		update["$pull"] = bson.M{"gameIds": bson.M{"$in": libraryRecord.WishlistIds}}

		if remove {
			update["$pull"] = bson.M{"wishlistIds": bson.M{"$in": libraryRecord.WishlistIds}}
		} else {
			update["$addToSet"] = bson.M{"wishlistIds": bson.M{"$each": libraryRecord.WishlistIds}}
		}
	}

	result, err := repo.db.Collection("library").UpdateOne(ctx, filter, update)
	return result, err
}

func (repo *libraryRepository) DeleteLibraryRecord(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	res, err := repo.db.Collection("library").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return &mongo.DeleteResult{}, err
	}
	return res, nil
}

func (repo *libraryRepository) DropLibraryRecords(ctx context.Context) {
	_, err := repo.db.Collection("library").DeleteMany(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
}

func (repo *libraryRepository) CreateIndices(ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "userId", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	name, err := repo.db.Collection("library").Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name of Index Created: " + name)
}
