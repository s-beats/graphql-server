package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-redis/redis/v8"
	"github.com/s-beats/graphql-todo/graph"
	"github.com/s-beats/graphql-todo/graph/generated"
	"github.com/s-beats/graphql-todo/graph/model"
	"github.com/s-beats/graphql-todo/queryservice"
	"github.com/s-beats/graphql-todo/usecase"
)

var graphQLServer *handler.Server

func GraphQLHandler(taskUsecase usecase.Task, userUsecase usecase.User, taskQueryService queryservice.Task, userQueryService queryservice.User) http.HandlerFunc {
	if graphQLServer == nil {
		r := &graph.Resolver{
			TaskUsecase:      taskUsecase,
			UserUsecase:      userUsecase,
			UserQueryService: userQueryService,
			TaskQueryService: taskQueryService,
			RedisClient:      redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
			TaskChannels:     make([]chan *model.TestSubscriptionPayload, 0),
		}
		// Contextはここで開始して良いのか？
		go func() {
			ctx := context.Background()
			pubsub := r.RedisClient.Subscribe(ctx, "pubsub_task")
			defer func() {
				pubsub.Close()
				log.Println("pubsub closed")
			}()
			log.Println("start subscribe")
			for {
				msg, err := pubsub.Receive(ctx)
				if err != nil {
					log.Println(err)
					continue
				}

				switch msg := msg.(type) {
				case *redis.Message:
					func() {
						r.TaskChannelsMutex.RLock()
						defer r.TaskChannelsMutex.RUnlock()
						log.Println("send channel", msg)
						for _, ch := range r.TaskChannels {
							ch <- &model.TestSubscriptionPayload{SubscriptionID: msg.Payload}
						}
					}()
				default:
					log.Printf("redis: unknown message: %T\n", msg)
				}
			}
		}()
		graphQLServer = handler.NewDefaultServer(generated.NewExecutableSchema(
			generated.Config{
				Resolvers: r,
			},
		))
		graphQLServer.AddTransport(&transport.Websocket{})
	}

	return func(w http.ResponseWriter, r *http.Request) {
		graphQLServer.ServeHTTP(w, r)
	}
}
