package entities

import "github.com/google/uuid"

type Movie struct {
	Id          string   `validate:"omitempty,uuid"`
	Title       string   `validate:"required"`
	Genre       []string `validate:"required"`
	Description string   `validate:"required"`
	Director    string   `validate:"required"`
	MainActors  []string `validate:"required"`
	Rating      float64  `validate:"required"`
}

func (m *Movie) SetId() {
	m.Id = uuid.New().String()
}
