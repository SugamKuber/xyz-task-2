package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"xyz-task-2/internals/configs"
	"xyz-task-2/internals/db"
	"xyz-task-2/internals/models"

	"github.com/gocql/gocql"
)

var (
	scyllaClient *db.ScyllaClient
	redisClient  *db.RedisClient
)

var (
	errorCategories    = []string{"Grammar", "Vocabulary", "Pronunciation", "Content"}
	errorSubcategories = map[string][]string{
		"Grammar":       {"Verb Agreement", "Tense", "Article Usage"},
		"Vocabulary":    {"Word Choice", "Idiomatic Expressions", "Collocations"},
		"Pronunciation": {"Stress", "Intonation", "Phoneme Production"},
		"Content":       {"Coherence", "Relevance", "Depth"},
	}
)

func init() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	scyllaClient, err = db.NewScyllaClient(cfg.ScyllaDB.ToScyllaConfig())
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB: %v", err)
	}

	redisClient, err = db.NewRedisClient(cfg.Redis.ToRedisConfig())
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	if err := resetAndPopulateData(); err != nil {
		log.Fatalf("Failed to reset and populate data: %v", err)
	}
}

func resetAndPopulateData() error {
	if err := dropTables(); err != nil {
		return fmt.Errorf("failed to drop tables: %v", err)
	}

	if err := createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}

	if err := populateData(); err != nil {
		return fmt.Errorf("failed to populate data: %v", err)
	}

	return nil
}

func dropTables() error {
	queries := []string{
		"DROP TABLE IF EXISTS users",
		"DROP TABLE IF EXISTS user_errors",
		"DROP TABLE IF EXISTS error_frequencies",
	}

	for _, query := range queries {
		if err := scyllaClient.Execute(query); err != nil {
			return err
		}
	}

	return nil
}

func createTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id UUID PRIMARY KEY,
            username TEXT
        )`,
		`CREATE TABLE IF NOT EXISTS user_errors (
		    user_id UUID,
		    conversation_id UUID,
		    timestamp TIMESTAMP,
		    error_category TEXT,
		    error_subcategory TEXT,
		    error_details TEXT,
		    PRIMARY KEY ((user_id), error_category, error_subcategory, conversation_id, timestamp)
		) WITH CLUSTERING ORDER BY (error_category ASC, error_subcategory ASC, conversation_id DESC, timestamp DESC);`,
		`CREATE TABLE IF NOT EXISTS error_frequencies (
		    user_id UUID,
		    error_category TEXT,
		    error_subcategory TEXT,
		    frequency counter,
		    PRIMARY KEY ((user_id), error_category, error_subcategory)
		);`,
	}
	// Improvements needed, this is a bad approach
	for _, query := range queries {
		if err := scyllaClient.Execute(query); err != nil {
			return err
		}
	}
	return nil
}
func generateAndInsertErrors(user models.User) error {
	for i := 0; i < 100; i++ {
		conversationID := gocql.TimeUUID()
		timestamp := time.Now().Add(-time.Duration(rand.Intn(30)) * 24 * time.Hour)

		errorCategory := errorCategories[rand.Intn(len(errorCategories))]
		errorSubcategory := errorSubcategories[errorCategory][rand.Intn(len(errorSubcategories[errorCategory]))]
		errorDetails := fmt.Sprintf("Error details for %s - %s", errorCategory, errorSubcategory)

		userErrorQuery := `
            INSERT INTO user_errors 
            (user_id, conversation_id, timestamp, error_category, error_subcategory, error_details) 
            VALUES (?, ?, ?, ?, ?, ?)
        `
		if err := scyllaClient.Execute(userErrorQuery, user.ID, conversationID, timestamp, errorCategory, errorSubcategory, errorDetails); err != nil {
			return err
		}

		updateFrequencyQuery := `
            UPDATE error_frequencies 
            SET frequency = frequency + 1 
            WHERE user_id = ? AND error_category = ? AND error_subcategory = ?
        `
		if err := scyllaClient.Execute(updateFrequencyQuery, user.ID, errorCategory, errorSubcategory); err != nil {
			return err
		}
	}
	// Improvements needed, but this doesnt matter since real time application doesnt have to mock

	fmt.Printf("Inserted data for user %s\n", user.Username)
	return nil
}

func populateData() error {

	users := generateUsers(5)
	for _, user := range users {
		if err := insertUser(user); err != nil {
			return err
		}
	}

	for _, user := range users {
		if err := generateAndInsertErrors(user); err != nil {
			return err
		}
	}

	return nil
}

func generateUsers(count int) []models.User {
	users := make([]models.User, count)
	for i := 0; i < count; i++ {
		users[i] = models.User{
			ID:       gocql.TimeUUID().String(),
			Username: fmt.Sprintf("user%d", i+1),
		}
	}
	return users
}

func insertUser(user models.User) error {
	query := "INSERT INTO users (id, username) VALUES (?, ?)"
	return scyllaClient.Execute(query, user.ID, user.Username)
}
