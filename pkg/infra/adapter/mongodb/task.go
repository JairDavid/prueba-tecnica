package adapter

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

	coll := t.conn.Collection("task")

	result, err := coll.InsertOne(context.Background(), task)
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
	var tasks []domain.Task

	coll := t.conn.Collection("task")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("[LOG] TaskRepository: ", err)
		return []domain.Task{}, err
	}

	if err = cursor.All(context.Background(), &tasks); err != nil {
		log.Println("[LOG] TaskRepository: ", err)
		return []domain.Task{}, err
	}

	return tasks, nil
}

// FindById implements port.ITaskRepository.
func (t TaskRepository) FindById(id string) (domain.Task, error) {
	var task domain.Task

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("[LOG] TaskRepository: ", err)
		return domain.Task{}, domain.TaskNotFound
	}

	coll := t.conn.Collection("task")
	err = coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&task)
	if err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("[LOG] TaskRepository: ", err)
			return domain.Task{}, domain.TaskNotFound
		}

		log.Println("[LOG] TaskRepository: ", err)
		return domain.Task{}, err
	}

	return task, nil
}

// UpdateById implements port.ITaskRepository.
func (t TaskRepository) UpdateById(id string, task domain.Task) (domain.Task, error) {
	panic("unimplemented")
}

// DeleteById implements port.ITaskRepository.
func (t TaskRepository) DeleteById(id string) (domain.Task, error) {
	panic("unimplemented")
}
