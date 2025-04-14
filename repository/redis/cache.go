package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheRepository struct {
	rdb *redis.Client
}

func NewCacheRepository(rdb *redis.Client) *CacheRepository {
	return &CacheRepository{rdb: rdb}
}

func (r *CacheRepository) SetToken(ctx context.Context, key, value string, ttl time.Duration) error {
	return r.rdb.Set(ctx, key, value, ttl).Err()
}

func (r *CacheRepository) GetToken(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}
