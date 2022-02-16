package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"auth-go/graph/generated"
	"auth-go/graph/model"
	"auth-go/handlers"
	"context"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (*model.AuthUser, error) {

	authUser, err := handlers.GetAuthUser(input.Email, input.Password)

	if err != nil {
		return authUser, gqlerror.Errorf(err.Error())
	}


	return authUser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//var users []*model.User
	//testUsers := model.User{
	//	ID:       "1",
	//	Name:     "Efrain",
	//	Email:    "efrain.be789@gmail.com",
	//	Password: "password123",
	//}
	//users = append(users, &testUsers)

	users := handlers.GetUsers()
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
