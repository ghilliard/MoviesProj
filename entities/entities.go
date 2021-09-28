package entities

import "github.com/google/uuid"

type Movie struct {
	Id          string
	Title       string
	Genre       []string
	Description string
	Director    string
	MainActors  []string
	Rating      float64
}

func (m *Movie) SetId() { //setting UUID
	m.Id = uuid.New().String()
}
