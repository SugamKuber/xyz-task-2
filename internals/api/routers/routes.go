package routers

import (
	"xyz-task-2/internals/api/handlers"
	"xyz-task-2/internals/db"
	"xyz-task-2/internals/middlewares"
	"xyz-task-2/internals/services/recommendation"

	"github.com/gorilla/mux"
)

func SetupRoutes(scyllaClient *db.ScyllaClient, redisClient *db.RedisClient) *mux.Router {
	router := mux.NewRouter()

	recommendationService := recommendation.NewService(scyllaClient, redisClient)

	exerciseHandler := handlers.NewExerciseHandler(recommendationService)
	healthHandler := handlers.NewHealthHandler()

	router.Use(middlewares.Logging)
	router.Use(middlewares.CORS)

	router.HandleFunc("/health", healthHandler.Check).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generate-exercise", exerciseHandler.GenerateExercise).Methods("GET")

	return router
}
