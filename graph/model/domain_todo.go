package model

import "gitlab.com/trivery-id/skadi/internal/todo/domain"

func NewTodo(in *domain.Todo) *Todo {
	out := Todo(*in)
	return &out
}

func NewTodos(in []domain.Todo) []Todo {
	var out []Todo
	for _, v := range in {
		out = append(out, Todo(v))
	}

	return out
}
