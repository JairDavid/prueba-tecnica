package provider

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"omnicloud.mx/tasks/pkg/app"
	adapter "omnicloud.mx/tasks/pkg/infra/adapter/mongodb"
	"omnicloud.mx/tasks/pkg/infra/api"
	"omnicloud.mx/tasks/pkg/infra/api/handler"
	"omnicloud.mx/tasks/pkg/infra/api/router"
	"omnicloud.mx/tasks/pkg/infra/resource"
)

type Container struct {
}

func New() *Container {
	return &Container{}
}

func (c *Container) Build() error {

	engine := chi.NewRouter()

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return fmt.Errorf("MONGO_URI environment var is emtpy")
	}

	port := os.Getenv("TASK_MICROSERVICE_PORT")
	if port == "" {
		return fmt.Errorf("TASK_MICROSERVICE_PORT environment var is emtpy")
	}

	client, err := resource.NewMongoDBClient(uri)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}

	//Dependency injection for task module
	taskRepository := adapter.NewTaskRepository(client.GetConnection())
	taskApp := app.NewTaskApp(taskRepository)
	taskHandler := handler.NewTaskHandler(taskApp)
	taskRouter := router.NewTaskRouter(taskHandler)

	//api server instance
	server := api.New(engine, taskRouter)
	server.MountRoutes()

	log.Println("[LOG] Server running on port: ", port)

	if err := http.ListenAndServe(":"+port, server.GetEngine()); err != nil {
		return fmt.Errorf("error starting server because: %w", err)
	}

	return nil
}
