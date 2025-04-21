package utils

import (
	"encoding/csv"
	"os"
)

func ParseCSV(filePath string) ([]map[string]string, error) {
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
	if len(records) < 1 {
		return nil, nil
	}

	headers := records[0]
	var data []map[string]string

	for _, row := range records[1:] {
		entry := make(map[string]string)
		for i, cell := range row {
			if i < len(headers) {
				entry[headers[i]] = cell
			}
		}
		data = append(data, entry)
	}

	return data, nil

}
