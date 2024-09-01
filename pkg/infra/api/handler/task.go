package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func (t TaskHandler) Save(w http.ResponseWriter, r *http.Request) {
	var taskDto domain.TaskDTO
	w.Header().Set("Content-Type", "application/json")

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

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(response.Common{
		Status:  http.StatusCreated,
		Message: "saved successfully",
		Data:    result,
	})
}

func (t TaskHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := t.taskApp.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(response.Error{
			Status:  http.StatusInternalServerError,
			Message: "something went wrong while quering data" + err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response.Common{
		Status:  http.StatusOK,
		Message: "consulted successfully",
		Data:    result,
	})
}

func (t TaskHandler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idDocument := chi.URLParam(r, "id")
	if idDocument == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response.Error{
			Status:  http.StatusBadRequest,
			Message: "id path-param is required",
		})
		return
	}

	result, err := t.taskApp.FindById(idDocument)
	if err != nil {

		if errors.Is(err, domain.TaskNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(response.Error{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response.Common{
		Status:  http.StatusOK,
		Message: "consulted successfully",
		Data:    result,
	})
}

func (t TaskHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (t TaskHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
