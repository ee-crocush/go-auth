package entity

import "time"

// User - сущность пользователя
type User struct {
	ID        string    // Уникальный идентификатор пользователя (например, UUID)
	Email     string    // Уникальный email
	Password  string    // Пароль пользователя
	CreatedAt time.Time // Дата и время создания пользователя
}

// Session - сущность сессии
type Session struct {
	ID        string    // Уникальный идентификатор сессии (например, UUID)
	UserID    string    // Уникальный идентификатор пользователя
	CreatedAt time.Time // Дата и время создания сессии
	// Другие поля сессии
}

//TODO определить, какие еще нужны сущности
