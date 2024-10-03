package recommendation

import (
	"encoding/json"
	"sort"
	"time"

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

func (s *Service) GetExerciseRecommendation(userID string) (models.ExerciseRecommendation, error) {
	cacheKey := "user:" + userID + ":exercise_recommendation"

	cachedData, err := s.redisClient.Get(cacheKey)
	if err == nil {
		var recommendation models.ExerciseRecommendation
		err = json.Unmarshal([]byte(cachedData), &recommendation)
		if err == nil {
			return recommendation, nil
		}
	}

	errors, err := s.scyllaClient.GetTopErrors(userID, 10)
	if err != nil {
		return models.ExerciseRecommendation{}, err
	}

	recommendation := models.ExerciseRecommendation{
		UserID:    userID,
		TopErrors: groupErrorsByCategory(errors),
	}

	jsonData, _ := json.Marshal(recommendation)
	s.redisClient.Set(cacheKey, jsonData, time.Hour)

	return recommendation, nil
}

func groupErrorsByCategory(errors []models.Error) []models.CategoryErrors {
	categoryMap := make(map[string][]models.Error)
	for _, err := range errors {
		categoryMap[err.Category] = append(categoryMap[err.Category], err)
	}

	var categoryErrors []models.CategoryErrors
	for category, errors := range categoryMap {
		categoryErrors = append(categoryErrors, models.CategoryErrors{
			Category: category,
			Errors:   errors,
		})
	}

	sort.Slice(categoryErrors, func(i, j int) bool {
		return categoryErrors[i].Errors[0].Frequency > categoryErrors[j].Errors[0].Frequency
	})

	return categoryErrors
}
