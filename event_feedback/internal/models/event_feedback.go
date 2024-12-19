package models

import "time"

type EventFeedback struct {
	ID        int       `json:"id"`
	EventID   int       `json:"event_id"`
	UserID    int       `json:"user_id"`
	Rating    int       `json:"rating"` // Пример: 1-5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
