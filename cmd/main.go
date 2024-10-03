package main

import (
	"log"
	"net/http"

	"xyz-task-2/internals/api/routers"
	"xyz-task-2/internals/configs"
	"xyz-task-2/internals/db"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	scyllaClient, err := db.NewScyllaClient(cfg.ScyllaDB.ToScyllaConfig())
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB: %v", err)
	}
	defer scyllaClient.Close()

	redisClient, err := db.NewRedisClient(cfg.Redis.ToRedisConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	router := routers.SetupRoutes(scyllaClient, redisClient)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
