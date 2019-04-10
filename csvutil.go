package csvutil

import (
	"encoding/csv"
	"io"
)

// CSVlines is a slice of slices that are columns from a csv
type CSVlines [][]string

// ReadCSV tranforms a csv reader intro a slice of lines
func ReadCSV(r io.Reader, comma, comment rune) (CSVlines, error) {
	var records CSVlines

	csvReader := csv.NewReader(r)
	csvReader.Comma = comma
	csvReader.Comment = comment

	// Read the first line and consume it (this is the colums header)
	_, err := csvReader.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}
