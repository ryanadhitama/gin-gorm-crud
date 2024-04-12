package repository

import (
	"gin-gorm-crud/model"
)

type TaskRepository interface {
	FindAll() ([]model.Task, error)
	FindById(taskId int) (task model.Task, err error)
	Save(task model.Task) error
	Update(task model.Task) error
	Delete(taskId int) error
}