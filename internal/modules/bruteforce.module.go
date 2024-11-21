package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type bruteForceModuleImpl struct{}

func NewBruteForceModule() Module {
	return &bruteForceModuleImpl{}
}

func (bfm *bruteForceModuleImpl) FindAvgFromFile(filename string, jobsNum int) error {
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

	sum := calculateSum(records, &err)

	if err != nil {
		return err
	}

	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("\nBrute Force module average: %f", sum/float64(len(records))))
	fmt.Println(fmt.Sprintf("Brute Force module read 1000 rows of CSV file took %s", elapsed))

	return nil
}
