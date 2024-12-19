package repository

import (
	"logs/internal/models"
	"github.com/jackc/pgx/v4"
	"context"

)

type LogRepository interface {
	CreateLog(log models.Log) (int, error)
	GetLogsByEventID(eventID int) ([]models.Log, error)
}

type logRepository struct {
	db *pgx.Conn
}

func NewLogRepository(db *pgx.Conn) LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) CreateLog(log models.Log) (int, error) {
	var id int
	err := r.db.QueryRow(context.Background(), "INSERT INTO logs (event_id, action, user_id, timestamp) VALUES ($1, $2, $3, $4) RETURNING id",
		log.EventID, log.Action, log.UserID, log.Timestamp).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *logRepository) GetLogsByEventID(eventID int) ([]models.Log, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, event_id, action, user_id, timestamp FROM logs WHERE event_id = $1", eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.Log
	for rows.Next() {
		var log models.Log
		if err := rows.Scan(&log.ID, &log.EventID, &log.Action, &log.UserID, &log.Timestamp); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
