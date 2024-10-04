package routers

import (
	"xyz-task-2/internals/api/handlers"
	"xyz-task-2/internals/db"
	"xyz-task-2/internals/middlewares"
	"xyz-task-2/internals/services/recommendation"
	"xyz-task-2/internals/services/users"

	"github.com/gorilla/mux"
)

func SetupRoutes(scyllaClient *db.ScyllaClient, redisClient *db.RedisClient) *mux.Router {
	router := mux.NewRouter()

	recommendationService := recommendation.NewService(scyllaClient, redisClient)
	userService := users.NewService(scyllaClient, redisClient)

	exerciseHandler := handlers.NewExerciseHandler(recommendationService)
	usersHandler := handlers.NewUserHandler(userService)
	healthHandler := handlers.NewHealthHandler()

	router.Use(middlewares.Logging)
	router.Use(middlewares.CORS)

	router.HandleFunc("/health", healthHandler.Check).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generate-exercise", exerciseHandler.GenerateExercise).Methods("GET")
	api.HandleFunc("/users", usersHandler.GetUsers).Methods("GET")

	return router
}
