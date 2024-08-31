package port

import "omnicloud.mx/tasks/pkg/domain"

type ITaskRepository interface {
	Save(task domain.Task) (domain.Task, error)
	FindAll() ([]domain.Task, error)
	FindById(id string) (domain.Task, error)
	UpdateById(id string, task domain.Task) (domain.Task, error)
	DeleteById(id string) (domain.Task, error)
}
