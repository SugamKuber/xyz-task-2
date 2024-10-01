package recommendation

import (
	"encoding/json"
	"time"

	"xyz-task-2/internals/database/redis"
	"xyz-task-2/internals/database/scylla"
	"xyz-task-2/internals/models"
)

type Service struct {
	scyllaClient *scylla.Client
	redisClient  *redis.Client
}

func NewService(scyllaClient *scylla.Client, redisClient *redis.Client) *Service {
	return &Service{
		scyllaClient: scyllaClient,
		redisClient:  redisClient,
	}
}

func (s *Service) GetTopErrors(userID string) (models.ExerciseRecommendation, error) {
	cacheKey := "user:" + userID + ":top_errors"

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
		UserID: userID,
		Errors: errors,
	}

	jsonData, _ := json.Marshal(recommendation)
	s.redisClient.Set(cacheKey, jsonData, time.Hour)

	return recommendation, nil
}

func (s *Service) GenerateExercise(userID string) (models.Exercise, error) {
	topErrors, err := s.GetTopErrors(userID)
	if err != nil {
		return models.Exercise{}, err
	}

	var exercise models.Exercise
	if len(topErrors.Errors) > 0 {
		topError := topErrors.Errors[0]
		exercise = models.Exercise{
			UserID:      userID,
			Category:    topError.Category,
			Subcategory: topError.Subcategory,
			Content:     generateExerciseContent(topError.Category, topError.Subcategory),
		}
	} else {
		exercise = models.Exercise{
			UserID:   userID,
			Category: "General",
			Content:  "Practice your general English skills.",
		}
	}

	return exercise, nil
}

func generateExerciseContent(category, subcategory string) string {

	return "Exercise content for " + category + " - " + subcategory
}
