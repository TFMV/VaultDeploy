package services

import (
	"encoding/csv"
	"os"
)

// ProcessCSV reads and parses a CSV file, returning the records as a slice of string slices
func ProcessCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
