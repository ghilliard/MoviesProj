package main

import (
	"MoviesProj/handlers"
	"MoviesProj/repo"
	"MoviesProj/service"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fn := "/Users/gabrielhilliard/Desktop/Pivot/MoviesProj/moviedb.json"

	ext := filepath.Ext(fn)
	if ext != ".json" {
		log.Fatal("File extension invalid")
	}

	//constructing each layer of the application by giving it the error beneath it
	r := repo.NewRepository(fn) //lowest layer
	svc := service.NewService(r)
	handler := handlers.NewMovieHandler(svc)
	router := handlers.ConfigureRouter(handler)

	svr := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatalln(svr.ListenAndServe())
}
