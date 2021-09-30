package resolver

import (
	"gitlab.com/trivery-id/skadi/internal/todo/services"
)

var TodoService *services.TodoService = &services.TodoService{}

// var TodoService ITodoService

// type ITodoService interface {
// 	CreateNewTodo(ctx context.Context, in services.CreateNewTodoInput) (*domain.Todo, error)
// 	GetAllTodos(ctx context.Context, userID uint64) ([]domain.Todo, error)
// }
