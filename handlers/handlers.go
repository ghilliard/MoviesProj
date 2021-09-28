package handlers

import (
	"MoviesProj/entities"
	"MoviesProj/repo"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		log.Fatal(err)
	}

	movieDb, err := repo.CallMovies(mv)
	if err != nil {
		log.Fatal(err)
	}

	//movieDb := repo.DataBase{}
	//movieDb.PostToDb(mv)

	movieBytes, err := json.MarshalIndent(movieDb, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("moviedb.json", movieBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(movieBytes)
}
