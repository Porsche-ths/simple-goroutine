package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type simpleModuleImpl struct{}

func NewSimpleModule() Module {
	return &simpleModuleImpl{}
}

func (sm *simpleModuleImpl) FindAvgFromFile(filename string, jobsNum int) error {
	start := time.Now()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	n := len(records) / jobsNum

	result := make([]float64, jobsNum)
	for i := range jobsNum {
		result[i] = calculateSum(records[i*n:(i+1)*n], &err)
	}

	if err != nil {
		return err
	}

	sum := 0.0
	for i := range jobsNum {
		sum += result[i]
	}

	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("\nSimple module average: %f", sum/float64(len(records))))
	fmt.Println(fmt.Sprintf("Simple module read 50M rows of CSV file took %s", elapsed))

	return nil
}
