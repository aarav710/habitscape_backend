package sessions

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
)

type SessionsService struct {
  Store *redisstore.RedisStore
}


func NewSessionsService(ctx context.Context, client *redis.Client) (*SessionsService, error) {
  store, err := redisstore.NewRedisStore(ctx, client)
	if err != nil {
    return nil, err
	}
	return &SessionsService{Store: store}, err
}