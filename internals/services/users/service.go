package users

import (
	// "encoding/json"
	// "time"

	"fmt"
	"xyz-task-2/internals/db"
	"xyz-task-2/internals/models"
)

type Service struct {
	scyllaClient *db.ScyllaClient
	redisClient  *db.RedisClient
}

func NewService(scyllaClient *db.ScyllaClient, redisClient *db.RedisClient) *Service {
	return &Service{
		scyllaClient: scyllaClient,
		redisClient:  redisClient,
	}
}

func (s *Service) GetUsers() ([]models.User, error) {
	// cacheKey := "allUsers"

	// cachedData, err := s.redisClient.Get(cacheKey)
	// if err == nil {
	// 	var users []models.User
	// 	err = json.Unmarshal([]byte(cachedData), &users)
	// 	if err == nil {
	// 		return users, nil
	// 	}
	// }

	users, err := s.scyllaClient.GetUsers()
	fmt.Println(users)
	if err != nil {
		return nil, err
	}
	// jsonData, _ := json.Marshal(users)
	// s.redisClient.Set(cacheKey, jsonData, time.Hour)

	return users, nil
}
