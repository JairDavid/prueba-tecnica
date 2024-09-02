package domain

import "errors"

var (
	TaskNotFound = errors.New("task not found")
)

// external struct
type TaskDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// domain struct
type Task struct {
	ID          string `bson:"_id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Completed   bool   `bson:"completed,omitempty"`
}

func ToTaskDTO(task Task) TaskDTO {
	return TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
}

func ToTask(taskDTO TaskDTO) Task {
	return Task{
		ID:          taskDTO.ID,
		Title:       taskDTO.Title,
		Description: taskDTO.Description,
		Completed:   taskDTO.Completed,
	}
}

func ToSliceTaskDTO(tasks []Task) []TaskDTO {
	slice := make([]TaskDTO, len(tasks))
	for i, item := range tasks {
		slice[i] = ToTaskDTO(item)
	}
	return slice
}
