package repository

import (
	"github.com/AituAbdiluly/todo-go"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID int, listID int) (todo.TodoList, error)
	DeleteByID(userID int, listID int) error
	UpdateByID(userID int, listID int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID int, itemID int) (todo.TodoItem, error)
	DeleteByID(userID int, itemID int) error
	UpdateByID(userID int, itemID int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPorstgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
