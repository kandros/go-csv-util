package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	csvutil "github.com/kandros/go-csv-util"
)

type movie struct {
	title    string
	director string
	year     int
}

func moviesToCSV(movies []movie, w io.Writer) error {
	writer := csv.NewWriter(w)
	writer.Comma = ','
	headers := []string{"movie title", "director", "year released"}
	var lines []csvutil.CSVline
	for _, m := range movies {
		lines = append(lines, m.toCSV())
	}
	err := csvutil.WriteCSV(headers, lines,
		w, ',')
	return err
}

func (m movie) toCSV() []string {
	return []string{m.title, m.director, strconv.Itoa(m.year)}
}

func main() {
	movies := []movie{
		movie{"Pulp Fiction", "Quentin Tarantino", 1994},
		movie{"Inception", "Christopher Nolan", 2010},
	}

	moviesToCSV(movies, os.Stdout)
}
