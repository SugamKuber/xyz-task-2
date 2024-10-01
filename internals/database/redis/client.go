package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

type Config struct {
	Address  string
	Password string
	DB       int
}

func NewClient(cfg Config) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(context.Background(), key, value, expiration).Err()
}

func (c *Client) Close() error {
	return c.client.Close()
}
