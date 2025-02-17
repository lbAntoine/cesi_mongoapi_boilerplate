package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDB(uri, dbName string, timeout int) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		client:   client,
		database: client.Database(dbName),
	}, nil
}

func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	return m.database.Collection(name)
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
