package entity

import "time"

// User - сущность пользователя
type User struct {
	ID        string    `json:"id"`         // Уникальный идентификатор пользователя (например, UUID)
	Email     string    `json:"email"`      // Уникальный email
	Password  string    `json:"password"`   // Пароль пользователя
	CreatedAt time.Time `json:"created_at"` // Дата и время создания пользователя
	IsActive  bool      `json:"is_active"`  // Флаг активности пользователя
}
