package handlers

import (
	"MoviesProj/entities"
	"MoviesProj/repo"
	//"MoviesProj/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// Instance of interface that defines the functionality that a repo must have in order to be called in this handlers layer
type Service interface {
	AddMovie(m entities.Movie) error
	ViewMovies() (repo.DataBase, error)
	FindMovieById (id string) (entities.Movie, error)
	DeleteMovieById (id string) error
	UpdateMovieById (id string, m entities.Movie) error
}

type MovieHandler struct {
	Svc Service
}

func NewMovieHandler(s Service) MovieHandler { //func that returns MovieHandler struct that is called in main
	return MovieHandler{
		Svc: s,
	}
}

func (mh MovieHandler) PostMovieHandler(w http.ResponseWriter, r *http.Request) {
	m := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = mh.Svc.AddMovie(m)
	if err != nil {
		switch err.Error() {
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //or StatusOK?
}

func (mh MovieHandler) GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	database, err := mh.Svc.ViewMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	movieBytes, err := json.MarshalIndent(database, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(movieBytes)
}

func (mh MovieHandler) GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //URL of the Request creates a map of route variables which are retrieved by this line
	id := vars["Id"]

	movieStruct, err := mh.Svc.FindMovieById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	movieBytes, err := json.MarshalIndent(movieStruct, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(movieBytes)
}

func (mh MovieHandler) DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := mh.Svc.DeleteMovieById(id)
	if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (mh MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	movieStruct := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&movieStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.UpdateMovieById(id, movieStruct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound) //or 500 error?
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
}
