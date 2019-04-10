package csvutil

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

// func TestWriteCSV(t *testing.T) {
// 	headers := []string{"cat name", "cat color"}
// 	var writer bytes.Buffer

// 	err := WriteCSV(headers, []CSVline{
// 		[]string{"micio", "black"},
// 		[]string{"micione", "orange"},
// 	},
// 		&writer, ',')
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	delimiter := []byte("\n")
// 	for {
// 		record, err := writer.Read(delimiter)
// 		if err == io.EOF {
// 			continue
// 		} else if err != nil {
// 			t.Fatal(err)
// 		}

// 		if string(record) != "ciaone" {
// 			t.Fatalf("\ngot:\n%v\nwant:\n %v", record, "ciaone")
// 		}
// 	}
// }

func TestReadCSV(t *testing.T) {
	csvFile, err := ioutil.ReadFile("./movies.csv")
	if err != nil {
		t.Error(err)
	}
	r := bytes.NewReader(csvFile)
	lines, err := ReadCSV(r, ';', '-')
	if err != nil {
		t.Error(err)
	}

	testTable := []CSVline{
		[]string{"Pulp Fiction", "Quentin Tarantino", "1994"},
		[]string{"Inception", "Christopher Nolan", "2010"},
	}

	if !reflect.DeepEqual(lines, testTable) {
		t.Fatalf("\ngot:\n%v\nwant:\n %v", lines, testTable)
	}

	if err != nil {
		t.Fatal(err)
	}
}
