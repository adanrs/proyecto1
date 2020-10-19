package main

/*
Adan Rivera Sanchez
18/10/2020
*/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	//"github.com/gorilla/mux"
)

func find(x string) int {
	for i, movie := range movies {
		if x == movie.Rank || x == movie.Title || x == movie.Description || x == movie.Director {
			return i
		}
	}
	return -1
}

/**
 * Method to Return the search
 *
 * @param GET, ID get the entity needed
 */
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	fmt.Println("Parameter:", id)
	if id == "movie" {
		w.Header().Set("Content-Type", "application/json")
		dataJson, _ := json.Marshal(movies)
		w.Write(dataJson)
	}
	checkError("Parse error", err)

	i := find(id)
	if i == -1 {
		fmt.Println("Parameter:", i)
		if id != "movie" {
			fmt.Fprintf(w, "No se encuentra la busqueda %v", id)
		}
		return
	}

	dataJson, err := json.Marshal(movies[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

/**
 * Method to Post to the CSV
 *
 * @param POST, handles the different post needed to fill the csv
 */
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	movie := Movie{}
	json.Unmarshal(body, &movie)
	movies = append(movies, movie)
	w.WriteHeader(200)
	return
}

/**
 * Method to update a entity with new data.
 *
 * @param PUT , handles the update of an entity
 */
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "application/json")
	id := path.Base(r.URL.Path)

	for index, item := range movies {
		if item.Rank == id {
			movies = append(movies[:index], movies[index+1:]...)
			var Movie Movie
			_ = json.NewDecoder(r.Body).Decode(&Movie)
			Movie.Rank = id
			movies = append(movies, Movie)
			json.NewEncoder(w).Encode(&Movie)
			return
		}
	}

	w.WriteHeader(200)
	return
}

/**
 * Method to delete
 *
 * @param DELETE, id entity to be deleted
 */
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {

	id := path.Base(r.URL.Path)
	i := find(id)
	if i == -1 {
		fmt.Println("ID Invalido")

	}
	movies = append(movies[:i], movies[i+1:]...)
	w.WriteHeader(200)
	return
}

func handleGetD(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	fmt.Println("Parameter:", id)
	if id == "movie" {
		w.Header().Set("Content-Type", "application/json")
		dataJson, _ := json.Marshal(movies)
		w.Write(dataJson)
	}
	checkError("Parse error", err)

	i := find(id)
	if i == -1 {
		fmt.Println("Parameter:", i)
		if id != "movie" {
			fmt.Fprintf(w, "No se encuentra el parametro %v", id)
		}
		return
	}

	dataJson, err := json.Marshal(movies[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}
