package service

import (
	"github.com/AituAbdiluly/todo-go"
	"github.com/AituAbdiluly/todo-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID int, listID int) (todo.TodoList, error)
	DeleteByID(userID int, listID int) error
	UpdateByID(userID int, listID int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userID int, listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID int, itemID int) (todo.TodoItem, error)
	DeleteByID(userID int, itemID int) error
	UpdateByID(userID int, itemID int, input todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
