package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	csvutil "github.com/kandros/go-csv-util"
)

type movie struct {
	title    string
	director string
	year     int
}

func csvToMovie(r io.Reader) ([]movie, error) {
	var movies []movie

	lines, err := csvutil.ReadCSV(r, ';', '-')
	if err != nil {
		return nil, err
	}

	for _, l := range lines {
		year, err := strconv.ParseInt(l[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := movie{l[0], l[1], int(year)}
		movies = append(movies, m)
	}

	return movies, nil
}

func main() {
	csvFile, err := ioutil.ReadFile("../../movies.csv")
	if err != nil {
		panic(err)
	}

	movies, err := csvToMovie(bytes.NewReader(csvFile))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", movies)
}
