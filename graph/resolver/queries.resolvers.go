package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/model"
)

func (r *queryResolver) GetTodos(ctx context.Context, userID *uint64) ([]model.Todo, error) {
	if userID == nil {
		return nil, fmt.Errorf("userID is required")
	}

	todos, err := TodoService.GetAllTodos(ctx, *userID)
	if err != nil {
		return nil, err
	}

	return model.NewTodos(todos), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
