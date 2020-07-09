package controller

import (
	"../repository"
	"../domain"
)

type TaskController struct {
	repository *repository.TaskRepository
}

func NewTaskController() (*TaskController, error) {
	r, err := repository.NewTaskRepository()
	if err != nil {
		return nil, err
	}
	return &TaskController{
		repository: r,
	}, nil
}

func (c *TaskController) Index() []domain.Task {
	return c.repository.GetAll()
}

