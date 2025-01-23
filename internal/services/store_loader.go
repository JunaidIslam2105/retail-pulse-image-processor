package services

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Store struct {
	Location  string
	StoreName string
	StoreID   string
}

var StoreMaster map[string]Store

func LoadStoreMaster(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening store master file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading CSV file: %v", err)
	}

	StoreMaster = make(map[string]Store)

	for _, record := range records[1:] {

		if len(record) < 3 {
			continue
		}

		store := Store{
			Location:  record[0],
			StoreName: record[1],
			StoreID:   record[2],
		}
		StoreMaster[store.StoreID] = store
	}
	return nil
}
