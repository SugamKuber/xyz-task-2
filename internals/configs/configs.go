package configs

import (
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


func NewConfig() *Config {
	return &Config{
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
