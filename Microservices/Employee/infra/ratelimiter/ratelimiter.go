package ratelimiter

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rhzyyn/Aozora/Microservices/Employee/config"
)

type Strategy interface {
	Allow(ctx context.Context, key string) bool
	HealthCheck(ctx context.Context) error
}

type RedisStrategy struct {
	client *redis.Client
	limit  int
	ttl    time.Duration
}

func NewRedisStrategy(cfg config.RateLimitConfig, client *redis.Client) *RedisStrategy {
	return &RedisStrategy{
		client: client,
		limit:  cfg.MaxRequestsPerMinute,
		ttl:    time.Minute,
	}
}

func (r *RedisStrategy) Allow(ctx context.Context, key string) bool {
	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return true // fail-open
	}
	if count == 1 {
		_ = r.client.Expire(ctx, key, r.ttl).Err()
	}
	return count <= int64(r.limit)
}

func (r *RedisStrategy) HealthCheck(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
