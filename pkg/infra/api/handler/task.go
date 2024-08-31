package handler

import (
	"net/http"

	"omnicloud.mx/tasks/pkg/app"
)

type ITaskHandler interface {
	Save(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	UpdateById(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
}

type TaskHandler struct {
	taskApp app.ITaskApp
}

func NewTaskHandler(taskApp app.ITaskApp) ITaskHandler {
	return TaskHandler{
		taskApp: taskApp,
	}
}

func (t TaskHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (t TaskHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (t TaskHandler) FindById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (t TaskHandler) Save(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (t TaskHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
