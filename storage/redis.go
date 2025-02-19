package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisLocker struct {
	redis *redis.Client
}

func NewRedisLocker(redis *redis.Client) *redisLocker {
	return &redisLocker{redis: redis}
}

func (d *redisLocker) LockSessionID(ctx context.Context, sessionID string) error {
	success, err := d.redis.SetNX(ctx, sessionID, "processing", time.Second*30).Result()
	if err != nil {
		return err
	}
	if !success {
		return fmt.Errorf("occupied session")
	}
	return nil
}
