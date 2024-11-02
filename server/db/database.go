package db

import (
	"context"
	"fmt"
	"log"

	"github.com/HemlockPham7/server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(config *config.EnvConfig) (*mongo.Collection, *mongo.Client) {
	MONGODB_URI := fmt.Sprintf(
		"mongodb+srv://%s:%s@cluster0.o0jwz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0",
		config.Username,
		config.Password,
	)

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")
	collection := client.Database(config.DatabaseName).Collection(config.CollectionName)

	return collection, client
}
