package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	"omnicloud.mx/tasks/pkg/domain"
	"omnicloud.mx/tasks/pkg/domain/port"
)

type TaskRepository struct {
	conn *mongo.Database
}

func NewTaskRepository(conn *mongo.Database) port.ITaskRepository {
	return TaskRepository{
		conn: conn,
	}
}

// Save implements port.ITaskRepository.
func (t TaskRepository) Save(task domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// FindAll implements port.ITaskRepository.
func (t TaskRepository) FindAll() ([]domain.Task, error) {
	panic("unimplemented")
}

// FindById implements port.ITaskRepository.
func (t TaskRepository) FindById(id string) (domain.Task, error) {
	panic("unimplemented")
}

// UpdateById implements port.ITaskRepository.
func (t TaskRepository) UpdateById(id string, task domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// DeleteById implements port.ITaskRepository.
func (t TaskRepository) DeleteById(id string) (domain.Task, error) {
	panic("unimplemented")
}
