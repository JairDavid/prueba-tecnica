package router

import (
	"github.com/go-chi/chi/v5"
	"omnicloud.mx/tasks/pkg/infra/api/handler"
)

type ITaskRouter interface {
	Mount() *chi.Mux
}

type TaskRouter struct {
	taskHandler handler.ITaskHandler
}

func NewTaskRouter(taskHandler handler.ITaskHandler) ITaskRouter {
	return TaskRouter{
		taskHandler: taskHandler,
	}
}

// TaskHttp implements ITaskRouter
func (t TaskRouter) Mount() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/tasks", t.taskHandler.FindAll)
	r.Get("/tasks/{id}", t.taskHandler.FindById)
	r.Post("/tasks", t.taskHandler.Save)
	r.Patch("/tasks/{id}", t.taskHandler.UpdateById)
	r.Delete("/tasks/{id}", t.taskHandler.DeleteById)

	return r
}
