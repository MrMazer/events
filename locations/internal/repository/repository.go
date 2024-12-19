package repository

import (
	"locations/internal/models"
	"github.com/jackc/pgx/v4"
	"context"
)

type LocationRepository interface {
	AddLocation(location models.Location) (int, error)
	GetLocationByID(locationID int) (models.Location, error)
}

type locationRepository struct {
	db *pgx.Conn
}

// Измените имя функции на NewLocationsRepository
func NewLocationsRepository(db *pgx.Conn) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) AddLocation(location models.Location) (int, error) {
	var id int
	err := r.db.QueryRow(context.Background(), "INSERT INTO locations (name, address, capacity, contact_person, phone_number) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		location.Name, location.Address, location.Capacity, location.ContactPerson, location.PhoneNumber).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *locationRepository) GetLocationByID(locationID int) (models.Location, error) {
	var location models.Location
	err := r.db.QueryRow(context.Background(), "SELECT id, name, address, capacity, contact_person, phone_number FROM locations WHERE id = $1", locationID).Scan(&location.ID, &location.Name, &location.Address, &location.Capacity, &location.ContactPerson, &location.PhoneNumber)
	if err != nil {
		return models.Location{}, err
	}
	return location, nil
}
