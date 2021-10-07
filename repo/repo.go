package repo

import (
	"MoviesProj/entities"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type DataBase struct {
	Movies []entities.Movie
}

type File struct {
	Filename string
}

func NewRepository(fn string) File {
	return File{
		Filename: fn,
	}
}

func (f File) CallMovie(mv entities.Movie) error { //method to File so we can call it in ReadFile
	db := DataBase{}
	jsonBytes, err := ioutil.ReadFile(f.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, &db)
	if err != nil {
		return err
	}
	db.Movies = append(db.Movies, mv)

	movieBytes, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(f.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f File) ViewMovies() (DataBase, error) {
	db := DataBase{}
	jsonBytes, err := ioutil.ReadFile(f.Filename)
	if err != nil {
		return db, err
	}
	err = json.Unmarshal(jsonBytes, &db)
	if err != nil {
		return db, err
	}
	return db, nil
}

func (f File) FindMovieById(id string) (entities.Movie, error) {
	db := DataBase{}
	jsonBytes, err := ioutil.ReadFile(f.Filename)
	if err != nil {
		return entities.Movie{}, err
	}
	err = json.Unmarshal(jsonBytes, &db)

	idFound := entities.Movie{}

	for _, v := range db.Movies {
		if v.Id == id {
			idFound = v
			return idFound, nil
		}
	}
	return idFound, nil
}

func (f File) DeleteMovieById(id string) error {
	db := DataBase{}
	jsonBytes, err := ioutil.ReadFile(f.Filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, &db)

	newDb := DataBase{} //new variable for Database that we use to append the ID and compare to original Database
						//so we can return an error saying not found

	dbLength := len(db.Movies)
	for _, v := range db.Movies {
		if v.Id == id {
			continue
		} else {
			newDb.Movies = append(newDb.Movies, v)
		}
	}

	if len(newDb.Movies) == dbLength {
		return errors.New("could not delete, movie not found")
	}
	movieBytes, err := json.MarshalIndent(newDb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(f.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}