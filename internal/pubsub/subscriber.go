package pubsub

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisSubscriber struct {
	rdb     *redis.Client
	channel string
	handler func(event JobEvent)
}

func (s *RedisSubscriber) Start(ctx context.Context) error {
	sub := s.rdb.Subscribe(ctx, s.channel)
	ch := sub.Channel()

	for {
		select {
		case msg := <-ch:
			var event JobEvent
			err := json.Unmarshal([]byte(msg.Payload), &event)
			if err != nil {
				log.Printf("Failed to unmarshal msg : %v", err)
			}
			s.handler(event)

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
