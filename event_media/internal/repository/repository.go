package repository

import (
	"event_media/internal/models"
	"github.com/jackc/pgx/v4"
	"context"

)

type EventMediaRepository interface {
	AddMedia(media models.EventMedia) (int, error)
	GetMediaByEventID(eventID int) ([]models.EventMedia, error)
}

type eventMediaRepository struct {
	db *pgx.Conn
}

func NewEventMediaRepository(db *pgx.Conn) EventMediaRepository {
	return &eventMediaRepository{db: db}
}

func (r *eventMediaRepository) AddMedia(media models.EventMedia) (int, error) {
	var id int
	err := r.db.QueryRow(context.Background(), "INSERT INTO event_media (event_id, media_type, media_url) VALUES ($1, $2, $3) RETURNING id",
		media.EventID, media.MediaType, media.MediaURL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *eventMediaRepository) GetMediaByEventID(eventID int) ([]models.EventMedia, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, event_id, media_type, media_url, uploaded_at FROM event_media WHERE event_id = $1", eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaItems []models.EventMedia
	for rows.Next() {
		var media models.EventMedia
		if err := rows.Scan(&media.ID, &media.EventID, &media.MediaType, &media.MediaURL, &media.UploadedAt); err != nil {
			return nil, err
		}
		mediaItems = append(mediaItems, media)
	}
	return mediaItems, nil
}
