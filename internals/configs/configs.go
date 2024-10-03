package configs

import (
	"errors"
	"xyz-task-2/internals/db"
)

type Config struct {
	ServerAddress string
	ScyllaDB      ScyllaDBConfig
	Redis         RedisConfig
}

type ScyllaDBConfig struct {
	Hosts    []string
	Keyspace string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func Load() (*Config, error) {
	config := &Config{
		ServerAddress: ":8080",
		ScyllaDB: ScyllaDBConfig{
			Hosts:    []string{"scylla:9042"},
			Keyspace: "system",
		},
		Redis: RedisConfig{
			Address:  "redis:6379",
			Password: "",
			DB:       0,
		},
	}

	if config.ServerAddress == "" {
		return nil, errors.New("server address cannot be empty")
	}
	if len(config.ScyllaDB.Hosts) == 0 {
		return nil, errors.New("no ScyllaDB hosts defined")
	}
	if config.Redis.Address == "" {
		return nil, errors.New("Redis address cannot be empty")
	}

	return config, nil
}

func (c *ScyllaDBConfig) ToScyllaConfig() db.ScyllaConfig {
	return db.ScyllaConfig{
		Hosts:    c.Hosts,
		Keyspace: c.Keyspace,
	}
}

func (c *RedisConfig) ToRedisConfig() db.RedisConfig {
	return db.RedisConfig{
		Address:  c.Address,
		Password: c.Password,
		DB:       c.DB,
	}
}
