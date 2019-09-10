package db

import (
	"context"
	"fmt"

	"github.com/nicolas2bert/ba-server/gen/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName = "ba"

func newMongoDB(dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:28083"))
	if err != nil {
		fmt.Printf("err %v", err)
	}
	err = client.Connect(nil)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	return client.Database(dbName)
}

var mongodb = newMongoDB(databaseName)

func CreateUser(ctx context.Context, user *models.User) error {
	filter := bson.M{"id": user.ID}
	upsert := true
	option := options.FindOneAndReplaceOptions{
		Upsert: &upsert,
	}
	res := mongodb.Collection("users").FindOneAndReplace(ctx, filter, *user, &option)
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		return res.Err()
	}
	return nil
}

func GetUser(ctx context.Context, id string) (*models.User, error) {
	filter := bson.M{"id": id}
	res := mongodb.Collection("users").FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var user = models.User{}
	err := res.Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
