package client

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"jwtgo/pkg/logging"
)

type MongodbClient struct {
	url    string
	logger *logging.Logger
}

func NewMongodbClient(url string, logger *logging.Logger) *MongodbClient {
	return &MongodbClient{
		url:    url,
		logger: logger,
	}
}

func (mc *MongodbClient) Connect() *mongo.Client {
	mc.logger.Info("Connecting to MongoDB...")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mc.url))
	if err != nil {
		mc.logger.Fatal("Error while connecting to MongoDB: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		mc.logger.Fatal("Error while pinging MongoDB: ", err)
	}

	return client
}
