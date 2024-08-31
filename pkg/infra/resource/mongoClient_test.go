package resource_test

import (
	"context"
	"os"
	"testing"
	"time"

	"omnicloud.mx/tasks/pkg/infra/resource"
)

func Test_mongodbConnection(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		t.Fatal("MONGO_URI enviroment var is emtpy")
	}

	client, err := resource.NewMongoDBClient("mongodb://admin:D6Rhr2ey7aMzuTK75gcS@localhost:27017/task-db?authSource=admin")
	if err != nil {
		t.Fatal("connection failed: ", err)
	}

	err = client.GetConnection().Client().Ping(ctx, nil)
	if err != nil {
		t.Log("connection success!")
	}

}
