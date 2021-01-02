package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// Client to connect to mongodb
var Client *mongo.Client

// Connect to create the connection to MongoDB
func Connect() {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//Cancel context to avoid memory leak
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	Client = client
	log.Printf("Connected to MongoDB!")
}
