package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/s-beats/graphql-todo/graph/generated"
	"github.com/s-beats/graphql-todo/graph/model"
	"github.com/s-beats/graphql-todo/infra/rdb"
	"github.com/samber/lo"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.CreateTaskPayload, error) {
	task, err := r.TaskUsecase.Create(ctx, input.Title, input.Text, input.UserID, input.Priority.String())
	if err != nil {
		return nil, err
	}
	if err := r.RedisClient.Publish(ctx, "pubsub_task", input.Text).Err(); err != nil {
		log.Println("failed to publish", err)
	}
	return &model.CreateTaskPayload{
		Task: &model.Task{
			ID:    task.ID,
			Title: task.Title,
			Text:  task.Text,
			// FIXME: priority
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.CreateUserPayload, error) {
	user, err := r.UserUsecase.Create(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	return &model.CreateUserPayload{
		User: &model.User{
			ID:   user.ID,
			Name: user.Name,
		},
	}, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, id *string, priority *model.TaskPriority) ([]*model.Task, error) {
	tasks, err := r.TaskQueryService.Search(ctx, rdb.SearchTasksParams{
		ID: lo.FromPtr(id),
	})
	if err != nil {
		return nil, err
	}
	return lo.Map(tasks, func(task *rdb.Task, _ int) *model.Task {
		return &model.Task{
			ID:    task.ID,
			Title: task.Title,
			Text:  task.Text,
			// FIXME: priority
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
	}), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// TestSubscription is the resolver for the TestSubscription field.
func (r *subscriptionResolver) TestSubscription(ctx context.Context, subscriptionID string) (<-chan *model.TestSubscriptionPayload, error) {
	ch := make(chan *model.TestSubscriptionPayload)
	r.TaskChannels = append(r.TaskChannels, ch)
	go func() {
		<-ctx.Done()
		log.Println("close subscription")
		r.TaskChannelsMutex.Lock()
		defer r.TaskChannelsMutex.Unlock()
		for i, c := range r.TaskChannels {
			if c == ch {
				r.TaskChannels = append(r.TaskChannels[:i], r.TaskChannels[i+1:]...)
				break
			}
		}
	}()
	return ch, nil
}

// Tasks is the resolver for the tasks field.
func (r *userResolver) Tasks(ctx context.Context, obj *model.User) ([]*model.Task, error) {
	tasks, err := r.TaskQueryService.Search(ctx, rdb.SearchTasksParams{
		UserID: obj.ID,
	})
	if err != nil {
		return nil, err
	}
	return lo.Map(tasks, func(task *rdb.Task, _ int) *model.Task {
		return &model.Task{
			ID:    task.ID,
			Title: task.Title,
			Text:  task.Text,
			// FIXME: priority
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
	}), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
