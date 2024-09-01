package app

import (
	"omnicloud.mx/tasks/pkg/domain"
	"omnicloud.mx/tasks/pkg/domain/port"
)

type ITaskApp interface {
	Save(taskDto domain.TaskDTO) (domain.TaskDTO, error)
	FindAll() ([]domain.TaskDTO, error)
	FindById(id string) (domain.TaskDTO, error)
	UpdateById(id string, taskDto domain.TaskDTO) (domain.TaskDTO, error)
	DeleteById(id string) (domain.TaskDTO, error)
}

type TaskApp struct {
	taskRepository port.ITaskRepository
}

func NewTaskApp(taskRepository port.ITaskRepository) ITaskApp {
	return TaskApp{
		taskRepository: taskRepository,
	}
}

// Save implements ITaskApp
func (t TaskApp) Save(taskdto domain.TaskDTO) (domain.TaskDTO, error) {
	task, err := t.taskRepository.Save(domain.ToTask(taskdto))
	if err != nil {
		return domain.TaskDTO{}, err
	}

	return domain.ToTaskDTO(task), nil
}

// FindAll implements ITaskApp
func (t TaskApp) FindAll() ([]domain.TaskDTO, error) {
	list, err := t.taskRepository.FindAll()
	if err != nil {
		return make([]domain.TaskDTO, 0), err
	}

	return domain.ToSliceTaskDTO(list), nil
}

// FindById implements ITaskApp
func (t TaskApp) FindById(id string) (domain.TaskDTO, error) {
	task, err := t.taskRepository.FindById(id)
	if err != nil {
		return domain.TaskDTO{}, err
	}
	return domain.ToTaskDTO(task), nil
}

// UpdateById implements ITaskApp
func (t TaskApp) UpdateById(id string, taskdto domain.TaskDTO) (domain.TaskDTO, error) {
	task, err := t.taskRepository.UpdateById(id, domain.ToTask(taskdto))
	if err != nil {
		return domain.TaskDTO{}, err
	}

	return domain.ToTaskDTO(task), nil
}

// DeleteById implements ITaskApp
func (t TaskApp) DeleteById(id string) (domain.TaskDTO, error) {
	task, err := t.taskRepository.DeleteById(id)
	if err != nil {
		return domain.TaskDTO{}, err
	}

	return domain.ToTaskDTO(task), nil
}
