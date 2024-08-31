package resource

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	conn *mongo.Database
}

func NewMongoDBClient(uri string) (*MongoDBClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dial, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error connection to MongoDB: %w", err)
	}

	return &MongoDBClient{
		conn: dial.Database("task-db"),
	}, nil
}

func (m *MongoDBClient) GetConnection() *mongo.Database {
	return m.conn
}

func (m *MongoDBClient) Close() error {
	if err := m.conn.Client().Disconnect(context.Background()); err != nil {
		return fmt.Errorf("error disconnecting from MongoDB: %w", err)
	}
	return nil
}
