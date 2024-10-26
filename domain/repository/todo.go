package repository

import (
	"aws-sns-local-go/domain"
	"aws-sns-local-go/domain/value"
)

type ToDoRepository interface {
	FindAll() ([]domain.ToDo, error)
	FindByID(todoId value.ToDoID) (*domain.ToDo, error)
	FindByUserID(userId value.UserID) ([]domain.ToDo, error)
	Store(todo domain.ToDo) error
	Update(todo domain.ToDo) error
	Delete(todoId value.ToDoID) error
}
