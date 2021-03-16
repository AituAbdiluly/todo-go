package service

import (
	"github.com/AituAbdiluly/todo-go"
	"github.com/AituAbdiluly/todo-go/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID int, listID int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

func (s *TodoListService) DeleteByID(userID int, listID int) error {
	return s.repo.DeleteByID(userID, listID)
}

func (s *TodoListService) UpdateByID(userID int, listID int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateByID(userID, listID, input)
}
