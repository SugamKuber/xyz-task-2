package db

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	redisClient *redis.Client
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func NewRedisClient(cfg RedisConfig) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{redisClient: client}, nil
}

func (c *RedisClient) Get(key string) (string, error) {
	return c.redisClient.Get(context.Background(), key).Result()
}

func (c *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return c.redisClient.Set(context.Background(), key, value, expiration).Err()
}

func (c *RedisClient) Close() error {
	return c.redisClient.Close()
}
