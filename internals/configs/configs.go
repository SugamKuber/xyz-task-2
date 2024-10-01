package configs

import (
	"errors"
	"xyz-task-2/internals/database/redis"
	"xyz-task-2/internals/database/scylla"
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
			Hosts:    []string{"scylla-node1:9042"},
			Keyspace: "stimuler_ai",
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


func (c *ScyllaDBConfig) ToScyllaConfig() scylla.Config {
	return scylla.Config{
		Hosts:    c.Hosts,
		Keyspace: c.Keyspace,
	}
}


func (c *RedisConfig) ToRedisConfig() redis.Config {
	return redis.Config{
		Address:  c.Address,
		Password: c.Password,
		DB:       c.DB,
	}
}
