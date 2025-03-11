package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}

type redisRepository struct {
	ctx context.Context
	rdb *redis.Client
}

func NewRedis(ctx context.Context, rdb *redis.Client) *redisRepository {
	return &redisRepository{
		ctx,
		rdb,
	}
}

func (r *redisRepository) Set(key string, value string, expiration time.Duration) error {
	err := r.rdb.Set(r.ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set key %s: %w", key, err)
	}

	return nil
}

func (r *redisRepository) Get(key string) (string, error) {
	data, err := r.rdb.Get(r.ctx, key).Result()

	switch {
	case err == redis.Nil:
		return "", nil // Kembalikan string kosong jika key tidak ditemukan
	case err != nil:
		return "", fmt.Errorf("failed to get key %s: %w", key, err) // Wrap error untuk informasi lebih jelas
	default:
		return data, nil
	}
}

func (r *redisRepository) Delete(key string) error {
	err := r.rdb.Del(r.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key %s: %w", key, err)
	}
	return nil
}
