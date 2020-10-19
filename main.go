package main

/*
Adan Rivera Sanchez
18/10/2020
*/
import (
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	var err error
	readData("IMDB-Movie-Data.csv")
	switch request.Method {
	case "GET":
		err = handleGet(writer, request)
	case "POST":
		err = handlePost(writer, request)
	case "PUT":
		err = handlePut(writer, request)
	case "DELETE":
		err = handleDelete(writer, request)
	}
	writeData("IMDB-Movie-Data.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func handlerD(writer http.ResponseWriter, request *http.Request) {
	var err error
	readData("IMDB-Movie-Data.csv")
	switch request.Method {
	case "GET":
		err = handleGetD(writer, request)

	}
	writeData("IMDB-Movie-Data.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func handlerA(writer http.ResponseWriter, request *http.Request) {
	var err error
	readData("IMDB-Actor-Data.csv")
	switch request.Method {
	case "GET":
		err = handleGetD(writer, request)

	}
	writeData("IMDB-Movie-Data.csv")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/movies/", handler)
	http.HandleFunc("/directors/", handlerD)
	http.HandleFunc("/actors/", handlerA)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
