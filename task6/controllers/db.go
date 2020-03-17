package controllers

import (
	"../models"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const (
	databaseName = "blog"
	tableName    = "posts"
)

func getPosts(db *mongo.Client) ([]models.Post, error) {
	c := db.Database(databaseName).Collection(tableName)

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, "Find")
	}

	res := make([]models.Post, 0, 1)

	if err := cur.All(context.Background(), &res); err != nil {
		return nil, errors.Wrap(err, "All")
	}

	return res, nil
}

func getPost(db *mongo.Client, id string) (models.Post, error) {
	c := db.Database(databaseName).Collection(tableName)
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errors.Wrap(err, "objToHex")
	}
	filter := bson.D{{Key: "_id", Value: docId}}
	res := c.FindOne(context.Background(), filter)

	post := new(models.Post)
	if err := res.Decode(post); err != nil {
		return models.Post{}, errors.Wrap(err, "decode")
	}

	return *post, err
}

func addPost(db *mongo.Client, post models.Post) error {
	c := db.Database(databaseName).Collection(tableName)
	_, err := c.InsertOne(context.Background(), post)

	return err
}

func editPost(db *mongo.Client, post *models.Post, id string) error {
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errors.Wrap(err, "objToHex")
	}
	filter := bson.D{{Key: "_id", Value: docId}}

	update := bson.D{{"$set",
		bson.D{
			{"title", post.Title},
			{"content", post.Content},
		},
	}}
	log.Println(update, filter)
	c := db.Database(databaseName).Collection(tableName)
	_, err = c.UpdateOne(context.Background(), filter, update)
	return err
}

func deletePost(db *mongo.Client, id string) error {
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errors.Wrap(err, "objToHex")
	}

	filter := bson.D{{Key: "_id", Value: docId}}

	c := db.Database(databaseName).Collection(tableName)
	_, err = c.DeleteOne(context.Background(), filter)

	return err
}
