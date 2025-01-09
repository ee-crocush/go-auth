package entity

import "time"

// Blacklist - модель для хранения заблокированных токенов
type Blacklist struct {
	ID          int       `json:"id"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
}
