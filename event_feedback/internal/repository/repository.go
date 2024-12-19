package repository

import (
	"event_feedback/internal/models"
	"github.com/jackc/pgx/v4"
	"context"
)

type EventFeedbackRepository interface {
	CreateFeedback(feedback models.EventFeedback) (int, error)
	GetFeedbackByEventID(eventID int) ([]models.EventFeedback, error)
}

type eventFeedbackRepository struct {
	db *pgx.Conn
}

func NewEventFeedbackRepository(db *pgx.Conn) EventFeedbackRepository {
	return &eventFeedbackRepository{db: db}
}

func (r *eventFeedbackRepository) CreateFeedback(feedback models.EventFeedback) (int, error) {
	var id int
	err := r.db.QueryRow(context.Background(), "INSERT INTO event_feedback (event_id, user_id, rating, comment) VALUES ($1, $2, $3, $4) RETURNING id",
		feedback.EventID, feedback.UserID, feedback.Rating, feedback.Comment).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *eventFeedbackRepository) GetFeedbackByEventID(eventID int) ([]models.EventFeedback, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, event_id, user_id, rating, comment, created_at FROM event_feedback WHERE event_id = $1", eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []models.EventFeedback
	for rows.Next() {
		var feedback models.EventFeedback
		if err := rows.Scan(&feedback.ID, &feedback.EventID, &feedback.UserID, &feedback.Rating, &feedback.Comment, &feedback.CreatedAt); err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, feedback)
	}
	return feedbacks, nil
}
