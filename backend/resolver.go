package backend

//go:generate go run github.com/99designs/gqlgen

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/edenriquez/graphql_loves_go/backend/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	todos []*models.Todo
	users []*models.User
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	todo := &models.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		User: models.User{
			Name: input.UserName,
			ID: input.UserID,
		},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewTodo) (*models.User, error) {
	user :=  &models.User{
			Name: input.UserName,
			ID: input.UserID,
		}
	r.users = append(r.users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }

// GET
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.users, nil
}


type todoResolver struct{ *Resolver }

// SET
func (r *todoResolver) Users(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return &models.User{ID: obj.User.ID, Name: "user " + obj.User.Name}, nil
}

