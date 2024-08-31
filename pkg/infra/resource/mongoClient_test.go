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

	client, err := resource.NewMongoDBClient(uri)
	if err != nil {
		t.Fatal("connection failed: ", err)
	}

	err = client.GetConnection().Client().Ping(ctx, nil)
	if err != nil {
		t.Log("connection success!")
	}

}
