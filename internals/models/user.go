package models

type ExerciseRecommendation struct {
	UserID    string  `json:"user_id"`
	TopErrors []Error `json:"topErrors"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
