package models

import "time"

type EventMedia struct {
	ID         int       `json:"id"`
	EventID    int       `json:"event_id"`
	MediaType  string    `json:"media_type"`  // Пример: "image", "video"
	MediaURL   string    `json:"media_url"`   // Ссылка на файл
	UploadedAt time.Time `json:"uploaded_at"`
}
