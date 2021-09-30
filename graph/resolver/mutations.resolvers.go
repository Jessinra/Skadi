package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/model"
	"gitlab.com/trivery-id/skadi/internal/todo/services"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodo) (*model.Todo, error) {
	todo, err := TodoService.CreateNewTodo(ctx, services.CreateNewTodoInput{
		Text:        input.Text,
		Description: input.Description,
		UserID:      input.UserID,
	})
	if err != nil {
		return nil, err
	}

	return model.NewTodo(todo), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
