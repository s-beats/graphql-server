package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/s-beats/graphql-todo/graph/generated"
	"github.com/s-beats/graphql-todo/graph/internal"
	"github.com/s-beats/graphql-todo/graph/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.CreateTaskPayload, error) {
	task, err := r.TaskUsecase.Create(ctx, input.Title, input.Text, input.UserID, input.Priority.String())
	if err != nil {
		return nil, err
	}

	return &model.CreateTaskPayload{
		Task: internal.ConvertTask(task),
	}, nil
}

func (r *queryResolver) Tasks(ctx context.Context, id *string, priority *model.TaskPriority) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
