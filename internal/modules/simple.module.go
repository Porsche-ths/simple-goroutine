package modules

import (
	"encoding/csv"
	"os"
	"strconv"
)

type SimpleModule struct {}

func NewSimpleModule() *SimpleModule {
	return &SimpleModule{}
}

func (sm *SimpleModule) FindAvgFromfile(filename string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for _, record := range records {
		value, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return 0, err
		}
		sum += value
	}

	return sum / float64(len(records)), nil
}
