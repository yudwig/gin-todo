package repository

import (
	"../domain"
	"../database"
)

type TaskRepository struct {
	db *database.MysqlDriver
}

func NewTaskRepository() (*TaskRepository, error) {
	db, err := database.NewDefaultMysqlDriver()
	if err != nil {
		return nil, err
	}
	return &TaskRepository{
		db: db,
	}, nil
}

func (r *TaskRepository) GetAll() []domain.Task {
	return nil
}

//func (r *TaskRepository)GetById(id int) domain.Task {
//}

func (r *TaskRepository) UpdateTitle(id int, title string) {

}

func (r *TaskRepository) UpdateCompleted(id int, completed int) {

}

func (r *TaskRepository) Delete(id int) {

}

//func (r *TaskRepository)Create(task *domain.Task) domain.Task {
//}


