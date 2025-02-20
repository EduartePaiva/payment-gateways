package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisLocker struct {
	redis *redis.Client
}

func NewRedisLocker(redis *redis.Client) *redisLocker {
	return &redisLocker{redis: redis}
}

func (d *redisLocker) LockSessionID(ctx context.Context, sessionID string) (bool, error) {
	return d.redis.SetNX(ctx, sessionID, "processing", time.Second*30).Result()
}

func (d *redisLocker) UnlockSessionID(ctx context.Context, sessionID string) error {
	_, err := d.redis.Del(ctx, sessionID).Result()
	return err
}
