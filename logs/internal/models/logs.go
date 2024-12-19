package models

import "time"

type Log struct {
	ID        int       `json:"id"`
	EventID   int       `json:"event_id"`
	Action    string    `json:"action"`    // Описание действия
	UserID    int       `json:"user_id"`   // Идентификатор пользователя, выполнившего действие
	Timestamp time.Time `json:"timestamp"` // Время совершения действия
}
