package storage

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisBehavior(t *testing.T) {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	assert.NoError(t, err)
	rdb := NewRedisLocker(redis.NewClient(opt))
	ctx := context.Background()

	defer rdb.DelSessionID(ctx, "coolID")

	ok, err := rdb.LockSessionID(ctx, "coolID")
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = rdb.LockSessionID(ctx, "coolID")
	assert.NoError(t, err)
	assert.False(t, ok)

}
