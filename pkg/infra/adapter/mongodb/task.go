package adapter

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := t.conn.Collection("task")

	result, err := coll.InsertOne(ctx, task)
	if err != nil {
		log.Println("[LOG] TaskRepository: ", err)
		return domain.Task{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		if !ok {
			log.Println("[LOG] TaskRepository: ", err)
			return domain.Task{}, fmt.Errorf("failed to convert object ID")
		}
		task.ID = oid.Hex()
	}

	return task, nil
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
