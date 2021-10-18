package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RmItem 는 id를 받아서, 해당 컬렉션에서 id가 일치하는 Item을 삭제한다.
func RmItem(client *mongo.Client, id string) error {
	collection := client.Database(*flagDBName).Collection("item")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
