package graph

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/s-beats/graphql-todo/graph/model"
	"github.com/s-beats/graphql-todo/queryservice"
	"github.com/s-beats/graphql-todo/usecase"
)

type Resolver struct {
	TaskUsecase usecase.Task
	UserUsecase usecase.User

	UserQueryService queryservice.User
	TaskQueryService queryservice.Task

	RedisClient       *redis.Client
	TaskChannels      []chan *model.TestSubscriptionPayload
	TaskChannelsMutex sync.RWMutex
}
