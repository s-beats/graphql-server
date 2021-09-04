package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/s-beats/graphql-server/graph/generated"
	"github.com/s-beats/graphql-server/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := make([]*model.Todo, 1, 1)
	rows, err := r.DB.Query("select * from todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var i int
	var (
		id        int
		title     string
		text      string
		userID    int
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)
	for rows.Next() {
		todos[i] = new(model.Todo)
		todos[i].User = new(model.User)
		err := rows.Scan(&id, &title, &text, &userID, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		log.Printf("get %#v", todos)
		todos[i].ID = strconv.Itoa(id)
		todos[i].Text = text
		todos[i].Done = true
		i++
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
