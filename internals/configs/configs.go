package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"xyz-task-2/internals/database/redis"
	"xyz-task-2/internals/database/scylla"
)

type Config struct {
	ServerAddress string         `yaml:"server_address"`
	ScyllaDB      ScyllaDBConfig `yaml:"scylla_db"`
	Redis         RedisConfig    `yaml:"redis"`
}

type ScyllaDBConfig struct {
	Hosts    []string `yaml:"hosts"`
	Keyspace string   `yaml:"keyspace"`
}

type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func Load() (*Config, error) {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
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
