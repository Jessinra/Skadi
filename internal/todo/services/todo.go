package services

import (
	"context"
	"time"

	"gitlab.com/trivery-id/skadi/internal/todo/domain"
)

type TodoService struct {
	todos  []domain.Todo
	currID uint64
}

var Service = &TodoService{
	todos: []domain.Todo{},
}

func (svc *TodoService) CreateNewTodo(_ context.Context, in CreateNewTodoInput) (*domain.Todo, error) {
	svc.currID++
	todo := domain.Todo{
		ID:          svc.currID,
		CreatedAt:   time.Now(),
		Text:        in.Text,
		UserID:      in.UserID,
		Description: in.Description,
	}

	svc.todos = append(svc.todos, todo)
	return &todo, nil
}

func (svc *TodoService) GetAllTodos(_ context.Context, userID uint64) ([]domain.Todo, error) {
	out := []domain.Todo{}
	for _, todo := range svc.todos {
		if todo.UserID == userID {
			out = append(out, todo)
		}
	}

	return out, nil
}
