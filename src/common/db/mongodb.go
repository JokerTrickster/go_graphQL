package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"main/src/common/aws"
	"time"
)

var (
	UserCollection        *mongo.Collection
	AddressBookCollection *mongo.Collection
	ProductCollection     *mongo.Collection
	ServiceCollection     *mongo.Collection
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func InitMongo() error {
	uri, err := aws.AwsGetParam("mongodb_url")
	if err != nil {
		return err
	}
	clientOptions := options.Client()
	clientOptions = clientOptions.ApplyURI(uri)
	clientOptions.SetMaxPoolSize(1)
	clientOptions.SetMinPoolSize(1)
	clientOptions.SetMaxConnIdleTime(20 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	MongoClient, _ = mongo.Connect(ctx, clientOptions)
	if err := PingMongo(); err != nil {
		return err
	}
	if err := InitCollection(); err != nil {
		return err
	}
	return nil
}

func InitCollection() error {
	mongo := MongoClient.Database("ryan_dev")
	UserCollection = mongo.Collection("user")
	return nil
}

func PingMongo() error {
	err := MongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("db 핑 통과")
	return nil
}
