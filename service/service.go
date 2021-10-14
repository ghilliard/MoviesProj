package service

import (
	"MoviesProj/entities"
	"MoviesProj/repo"
	"errors"
	"github.com/google/uuid"
)

// Instance of interface that defines the functionality that a repo struct must have in order to be called in this service layer
type Repository interface {
	AddMovie(m entities.Movie) error
	ViewMovies() (repo.DataBase, error)
	FindMovieById (id string) (entities.Movie, error)
	DeleteMovieById (id string) error
	UpdateMovieById (id string, m entities.Movie) error
}

type Service struct { //create a struct of repo.Repo so you can use it as a method
	Repo Repository
}

//constructor func that that accepts interface and returns Service struct that is called in main
func NewService(r Repository) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) AddMovie(m entities.Movie) error {
	m.Id = uuid.New().String()

	if m.Rating >= 0 && m.Rating <= 10 {
		err := s.Repo.AddMovie(m)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid rating")
}

func (s Service) ViewMovies() (repo.DataBase, error) {
	db, err := s.Repo.ViewMovies()
	if err != nil {
		return db, err
	}
	return db, nil
}

func (s Service) FindMovieById(id string) (entities.Movie, error) {
	movie, err := s.Repo.FindMovieById(id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}

func (s Service) DeleteMovieById(id string) error {
	err := s.Repo.DeleteMovieById(id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) UpdateMovieById(id string, m entities.Movie) error {
	if id != m.Id {
		return errors.New("id must match url id")
	}

	err := s.Repo.UpdateMovieById(id, m)
	if err != nil {
		return err
	}
	return nil
}