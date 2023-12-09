package redispkg

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisRepo interface {
	SetValue(ctx context.Context, key string, field string, value interface{}) error
	GetValue(ctx context.Context, key, field string) (string, error)
	GetAll(ctx context.Context, key string) (map[string]string, error)
	HDelete(ctx context.Context, key string, fields ...string) error
	FlushRedis(ctx context.Context) error
}

func NewClent(cfg *Redis) RedisRepo {
	cfg.rClient = redis.NewClient(&redis.Options{
		Addr:     cfg.host,
		Username: cfg.useName,
		Password: cfg.password,
		Protocol: cfg.protocol,
	})
	return cfg
}

func (r *Redis) SetValue(ctx context.Context, key string, field string, value interface{}) error {
	return r.rClient.HSet(ctx, key, field, value).Err()
}

func (r *Redis) GetValue(ctx context.Context, key, field string) (string, error) {
	return r.rClient.HGet(ctx, key, field).Result()
}

func (r *Redis) GetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.rClient.HGetAll(ctx, key).Result()
}

func (r *Redis) HDelete(ctx context.Context, key string, fields ...string) error {
	return r.rClient.HDel(ctx, key, fields...).Err()
}

func (r *Redis) FlushRedis(ctx context.Context) error {
	return r.rClient.FlushAll(ctx).Err()
}
