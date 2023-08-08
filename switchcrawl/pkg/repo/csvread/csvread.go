package csvread

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/tungtanium/go-course23/switchcrawl/pkg/model"
)

// Read switch data from saved .csv file to set up database
func CsvReader(filepath string) []model.SwitchData {
	var out []model.SwitchData

	csvfile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Data from: ", filepath)

	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		s := model.SwitchData{
			Name:        record[0],
			Type:        record[1],
			Actuation:   record[2],
			PreTravel:   record[3],
			TotalTravel: record[4],
		}
		out = append(out, s)
	}
	return out
}
