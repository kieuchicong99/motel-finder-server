package infrastructure

import (
	"context"
	"fmt"
	"server/utilities"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoGoDriverConnector struct {
	client *mongo.Client
}

func (c *MongoGoDriverConnector)Connect()(*mongo.Client, error)  {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin%40123@cluster0.ce2mm.mongodb.net/motelfinder?retryWrites=true&w=majority")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		utilities.ErrLog.Printf("Connect error, err: %v", err)
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		utilities.ErrLog.Printf("Ping error, err: %v", err)
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil
}
