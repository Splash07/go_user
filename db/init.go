package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Splash07/go_user/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

const (
	DbName  = "User"
	ColName = "user"
)

func Test() interface{} {
	//
	fmt.Println("connect & insert db")
	return insertNumber()
}

func init() {
	connect()
}

func insertNumber() interface{} {
	collection := Client.Database("testing").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		return nil
	}
	id := res.InsertedID
	return id
}

func connect() {
	log.Println("Try to connect mongo:", config.Config.Mongo.Uri)
	clientOptions := options.Client().ApplyURI(config.Config.Mongo.Uri)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(4)
	clientOptions.SetReadPreference(readpref.Nearest())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping error: %v", err)
	}
	Client = client
}
