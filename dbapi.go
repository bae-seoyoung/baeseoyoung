package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddItem 은 데이터베이스에 Item(Project, Card, Client)를 넣는다.
func AddItem(client *mongo.Client, i Item) error {
	collection := client.Database(*flagDBName).Collection("item")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, i)
	if err != nil {
		return err
	}
	return nil
}

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
