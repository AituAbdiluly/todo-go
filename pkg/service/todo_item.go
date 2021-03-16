package service

import (
	"github.com/AituAbdiluly/todo-go"
	"github.com/AituAbdiluly/todo-go/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userID int, listID int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		// if list doesn't exist and does not belong to user
		return 0, err
	}

	return s.repo.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userID, listID)
}

func (s *TodoItemService) GetByID(userID int, itemID int) (todo.TodoItem, error) {
	return s.repo.GetByID(userID, itemID)
}

func (s *TodoItemService) DeleteByID(userID int, itemID int) error {
	return s.repo.DeleteByID(userID, itemID)
}

func (s *TodoItemService) UpdateByID(userID int, itemID int, input todo.UpdateItemInput) error {
	return s.repo.UpdateByID(userID, itemID, input)
}
