package entity

import "time"

// Session - сущность сессии
type Session struct {
	ID        string    // Уникальный идентификатор сессии (например, UUID)
	UserID    string    // Уникальный идентификатор пользователя
	CreatedAt time.Time // Дата и время создания сессии
	// Другие поля сессии
}
