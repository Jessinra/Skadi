package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/model"
	todoSvc "gitlab.com/trivery-id/skadi/internal/todo/services"
	userSvc "gitlab.com/trivery-id/skadi/internal/user/services"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.User, error) {
	user, err := userService.RegisterUser(ctx, userSvc.RegisterUserInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	user, err := userService.UpdateUser(ctx, userSvc.UpdateUserInput{
		ID:                input.ID,
		Name:              input.Name,
		PhoneNumber:       input.PhoneNumber,
		ProfilePictureURL: input.ProfilePictureURL,
		CurrencyMain:      input.CurrencyMain,
		CurrencySub:       input.CurrencySub,
	})
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *mutationResolver) UpdateUserPassword(ctx context.Context, input model.UpdateUserPassword) (bool, error) {
	if err := userService.UpdateUserPassword(ctx, userSvc.UpdateUserPasswordInput{
		ID:          input.ID,
		Password:    input.Password,
		NewPassword: input.NewPassword,
	}); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodo) (*model.Todo, error) {
	todo, err := todoService.CreateNewTodo(ctx, todoSvc.CreateNewTodoInput{
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
