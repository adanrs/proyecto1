package main

/*
Adan Rivera Sanchez
18/10/2020
*/
import (
	"encoding/csv"
	"os"
)

type Director struct {
	Rank     string `json:"rank"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Actors   string `json:"actors"`
}

var directors []Director

func readDataD(filePath string) {
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	directors = []Director{}

	for _, record := range records {
		Director := Director{
			Rank:     record[0],
			Title:    record[1],
			Director: record[3],
			Actors:   record[4],
		}

		directors = append(directors, Director)
	}
	file.Close()
}

func writeDataD(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, Director := range directors {
		record := []string{
			Director.Rank,
			Director.Title,
			Director.Director,
			Director.Actors}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}
	writer.Flush()
	file.Close()
}
