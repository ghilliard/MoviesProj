package handlers

import (
	"MoviesProj/entities"
	"MoviesProj/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MovieHandler struct {
	serv service.Service
}

func NewMovieHandler(s service.Service) MovieHandler { //func that returns MovieHandler struct that is called in main
	return MovieHandler{
		serv: s,
	}
}

func (mh MovieHandler) PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}
	err = mh.serv.CallMovie(mv)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //or StatusOK?
}

func (mh MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movDb, err := mh.serv.ViewMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	movieDb, err := json.MarshalIndent(movDb, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(movieDb)
}

func (mh MovieHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //create a map of route variables which are retrieved by this line
	id := vars["Id"]

	mvId, err := mh.serv.FindMovieById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	movie, err := json.MarshalIndent(mvId, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(movie)
}

func (mh MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := mh.serv.DeleteMovieById(id)
	if err != nil {
		switch err.Error() {
		case "could not delete, movie not found":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}