package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(mh MovieHandler) *mux.Router { //func that is called in main
	r := mux.NewRouter()

	r.HandleFunc("/movie", mh.PostMovieHandler).Methods("POST")
	r.HandleFunc("/movie", mh.GetMoviesHandler).Methods("GET")
	r.HandleFunc("/movie/{Id}", mh.GetByIdHandler).Methods("GET")
	r.HandleFunc("/movie/{Id}", mh.DeleteMovieHandler).Methods("DELETE")
	r.HandleFunc("/movie/{Id}", mh.UpdateMovieHandler).Methods("PUT")

	return r
}
