package repo

import (
	"MoviesProj/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DataBase struct {
	Movies []entities.Movie `json:"Movies"`
	//MoviesMap map[string] entities.Movie
}

//func (db *DataBase) PostToDb (movie entities.Movie) {
//	db.Movies = append(db.Movies, movie)
//}
//function that unmarshals
func CallMovies(mv entities.Movie) (DataBase, error) {
	db := DataBase{}
	jsonBytes, err := ioutil.ReadFile("jsonFile/moviedb.json")
	if err != nil {
		return db, err
	}
	mv.SetId()
	err = json.Unmarshal(jsonBytes, &db)
	if err != nil {
		return db, err
	}
	db.Movies = append(db.Movies, mv)

	movieBytes, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("jsonFile/moviedb.json", movieBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return db, nil
}

//function that marshals
//func (db DataBase) CreateMovie(movie entities.Movie) (entities.Movie, error) {
//	repo, _ := CallMovies()
//	repo.Movies = append(repo.Movies, movie)
//	sendFile, err := json.Marshal(movie)
//	jsonFile, err := json.MarshalIndent(repo, "", " ")
//	if err != nil {
//		return movie, err
//	}
//	err = ioutil.WriteFile("jsonFile/moviedb.json", jsonFile, 0644)
//	if err != nil {
//		return entities.Movie{}, err
//	}
//	return sendFile, err
//}
//
//
//func (db DataBase) FindMovieById(Id string) entities.Movie {
//	repo, _ := CallMovies()
//	//movie := repo.MoviesMap[Id]
//
//	return movie
//
//}

//func (db DataBase) DeleteMovie
