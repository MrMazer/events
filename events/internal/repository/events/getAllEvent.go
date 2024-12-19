package events

import (
	"context"
	"events/internal/models"
)

const (
	//getAllEventsQuery = `
	//   SELECT id, name_event, shape, place, begin_time, duration
	//   FROM events
	//   WHERE user_id = $1
	//
	//`
	getAllEventsQuery = `
        SELECT e.id, e.name_event, e.shape, e.place, e.begin_time, e.duration
        FROM events e
        FULL JOIN participation p ON e.id = p.event_id
        WHERE p.user_id = $1
    `
)

func (r *eventRepository) GetAllEvents(userID int64) ([]models.Event, error) {

	rows, err := r.conn.Query(context.Background(), getAllEventsQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err = rows.Scan(&event.ID, &event.NameEvent, &event.Shape, &event.Place, &event.BeginTime, &event.Duration)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
