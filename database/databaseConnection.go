package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var myEnv map[string]string
var MONGODB_URL string = "mongodb+srv://quocky:quocky@golang-restaurant-manag.lni0yjt.mongodb.net/"

var DB_NAME = "golang-restaurant-management"

func DBinstance() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_URL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb!")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(DB_NAME).Collection(collectionName)
	return collection
}
