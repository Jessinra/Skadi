package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id uint64) (*model.User, error) {
	user, err := userService.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *queryResolver) GetTodos(ctx context.Context, userID *uint64) ([]model.Todo, error) {
	if userID == nil {
		return nil, errors.New("userID is required")
	}

	todos, err := todoService.GetAllTodos(ctx, *userID)
	if err != nil {
		return nil, err
	}

	return model.NewTodos(todos), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
