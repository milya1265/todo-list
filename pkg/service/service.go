package service

import (
	todo_app "github.com/milya1265/todo-list.git"
	"github.com/milya1265/todo-list.git/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
