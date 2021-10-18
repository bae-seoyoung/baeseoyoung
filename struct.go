package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Author     string             `json:"author" bson:"author"`
	Createtime string             `json:"createtime" bson:"createtime"`
	Edittime   string             `json:"edittime" bson:"edittime"`
	Title      string             `json:"title" bson:"title"`
	Tags       []string           `json:"tags" bson:"tags"`
	Hide       bool               `json:"hide" bson:"hide"`
	Contents   []Content          `json:"contents" bson:"contents"`
}

type Content struct {
	Type string `json:"type" bson:"type"`
	Body string `json:"body" bson:"body"`
}
