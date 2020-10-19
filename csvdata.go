package main

/*
Adan Rivera Sanchez
18/10/2020
*/
import (
	"encoding/csv"
	"log"
	"os"
)

type Movie struct {
	Rank        string `json:"rank"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Actors      string `json:"actors"`
	Year        string `json:"year"`
	Runtime     string `json:"runtime"`
	Rating      string `json:"rating"`
	Votes       string `json:"votes"`
	Revenue     string `json:"revenue"`
	Metascore   string `json:"metascore"`
}

var movies []Movie

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func readData(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	movies = []Movie{}

	for _, record := range records {
		movie := Movie{
			Rank:        record[0],
			Title:       record[1],
			Genre:       record[2],
			Description: record[3],
			Director:    record[4],
			Actors:      record[5],
			Year:        record[6],
			Runtime:     record[7],
			Rating:      record[8],
			Votes:       record[9],
			Revenue:     record[10],
			Metascore:   record[11]}

		movies = append(movies, movie)
	}
	file.Close()
}

func writeData(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, movie := range movies {
		record := []string{
			movie.Rank,
			movie.Title,
			movie.Genre,
			movie.Description,
			movie.Director,
			movie.Actors,
			movie.Year,
			movie.Runtime,
			movie.Rating,
			movie.Votes,
			movie.Revenue,
			movie.Metascore}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
