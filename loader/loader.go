package loader

import (
	"encoding/csv"
	"io"
	"log"
)

type CustomerData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func LoadData(r io.Reader) *[]*CustomerData {
	reader := csv.NewReader(r)

	ret := make([]*CustomerData, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		if err != nil {
			log.Println(err)
		}
		customer := &CustomerData{
			ID:   row[0],
			Name: row[1],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, customer)
	}
	return &ret
}
