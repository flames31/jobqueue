package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisPublisher struct {
	rdb     *redis.Client
	channel string
}

func NewRedisPublisher(addr, password string, db int, channel string) *RedisPublisher {
	rds := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisPublisher{
		rdb:     rds,
		channel: channel,
	}
}

func (r *RedisPublisher) Publish(event JobEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error marshalling event into json : %w", err)
	}

	return r.rdb.Publish(context.Background(), r.channel, data).Err()
}
