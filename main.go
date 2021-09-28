package main

import (
	"MoviesProj/handlers"
	"log"
)

func main() {
	svr := handlers.NewServer()
	log.Fatal(svr.ListenAndServe())
}
