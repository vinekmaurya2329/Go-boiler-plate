package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectDB initializes the MongoDB connection
func ConnectDB() *mongo.Client {

	MONGO_URI := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(MONGO_URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("MongoDB connection error:")
		log.Fatal("MongoDB connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	fmt.Println(" Connected to MongoDB!")
	Client = client
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("goDb").Collection(collectionName)
}
