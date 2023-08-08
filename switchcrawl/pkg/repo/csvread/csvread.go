package csvread

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Read switch data from saved .csv file to set up database
func CsvReader(filepath string) {

	csvfile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Reading from: ", filepath)

	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Type %T %v\n", record, record[0])
	}
}
