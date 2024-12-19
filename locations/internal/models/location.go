package models

type Location struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Capacity      int    `json:"capacity"`
	ContactPerson string `json:"contact_person"`
	PhoneNumber   string `json:"phone_number"`
}
