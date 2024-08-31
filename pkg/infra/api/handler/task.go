package handler

import (
	"encoding/json"
	"net/http"

	"omnicloud.mx/tasks/pkg/app"
	"omnicloud.mx/tasks/pkg/domain"
	"omnicloud.mx/tasks/pkg/infra/api/response"
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
	var taskDto domain.TaskDTO

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewDecoder(r.Body).Decode(&taskDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.Error{
			Status:  http.StatusBadRequest,
			Message: "something went wrong while serializing data: " + err.Error(),
		})
		return
	}

	result, err := t.taskApp.Save(taskDto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.Error{
			Status:  http.StatusInternalServerError,
			Message: "something went wrong while saving data" + err.Error(),
		})
		return
	}

	_ = json.NewEncoder(w).Encode(response.Common{
		Status:  http.StatusCreated,
		Message: "saved successfully",
		Data:    result,
	})
}

func (t TaskHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
