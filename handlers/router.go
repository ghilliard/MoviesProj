package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(mh MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", mh.PostNewMovie).Methods("POST")
	r.HandleFunc("/movie", mh.GetAllMovies).Methods("GET")
	r.HandleFunc("/movie/{Id}", mh.GetById).Methods("GET")
	r.HandleFunc("/movie/{Id}", mh.DeleteMovie).Methods("DELETE")

	return r
}
