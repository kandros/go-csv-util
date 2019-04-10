package csvutil

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestCsvUtil(t *testing.T) {
	csvFile, err := ioutil.ReadFile("./movies.csv")
	if err != nil {
		t.Error(err)
	}
	r := bytes.NewReader(csvFile)
	lines, err := ReadCSV(r, ';', '-')
	if err != nil {
		t.Error(err)
	}

	testTable := CSVlines{
		[]string{"Pulp Fiction", "Quentin Tarantino", "1994"},
		[]string{"Inception", "Christopher Nolan", "2010"},
	}

	if !reflect.DeepEqual(lines, testTable) {
		t.Fatalf("\ngot:\n%v\nwant:\n %v", lines, testTable)
	}

	if err != nil {
		panic(err)
	}
}
