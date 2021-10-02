package service

import (
	"MoviesProj/entities"
	"MoviesProj/repo"
	"github.com/google/uuid"
)

type Service struct { //create a struct of repo.File so you can use it as a method
	File repo.File
}

func NewService(f repo.File) Service {
	return Service{
		File: f,
	}
}

func (s Service) CallMovie(mv entities.Movie) error {
	mv.Id = uuid.New().String()

	err := s.File.CallMovie(mv)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) ViewMovies() (repo.DataBase, error) {
	db, err := s.File.ViewMovies()
	if err != nil {
		return db, err
	}
	return db, nil
}

func (s Service) FindMovieById(id string) (entities.Movie, error) {
	movie, err := s.File.FindMovieById(id)
	if err != nil {
		return movie, err
	}
	return movie, nil
}
