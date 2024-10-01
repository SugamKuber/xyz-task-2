package models

type ExerciseRecommendation struct {
	UserID string  `json:"user_id"`
	Errors []Error `json:"errors"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type Exercise struct {
	UserID      string `json:"user_id"`
	Category    string `json:"category"`
	Subcategory string `json:"subcategory"`
	Content     string `json:"content"`
}

type Conversation struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Timestamp int64  `json:"timestamp"`
}

type Utterance struct {
	ID             string `json:"id"`
	ConversationID string `json:"conversation_id"`
	UserID         string `json:"user_id"`
	Content        string `json:"content"`
	Timestamp      int64  `json:"timestamp"`
}
