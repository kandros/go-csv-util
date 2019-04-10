package csvutil

import (
	"encoding/csv"
	"io"
)

// CSVlines is a slice of slices that are columns from a csv
type CSVline []string

// ReadCSV tranforms a csv reader intro a slice of lines
func ReadCSV(r io.Reader, commaSymbol, commentSymbol rune) ([]CSVline, error) {
	var records []CSVline

	csvReader := csv.NewReader(r)
	csvReader.Comma = commaSymbol
	csvReader.Comment = commentSymbol

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

// WriteCSV write csv to a Writer
func WriteCSV(headers []string, lines []CSVline, w io.Writer, commaSymbol rune) error {
	writer := csv.NewWriter(w)
	writer.Comma = commaSymbol

	// add columns headers
	err := writer.Write(headers)
	if err != nil {
		return err
	}

	for _, l := range lines {
		err := writer.Write(l)
		if err != nil {
			return err
		}
	}

	writer.Flush()

	return nil
}
