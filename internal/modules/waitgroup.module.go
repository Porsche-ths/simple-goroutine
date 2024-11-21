package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type waitGroupModuleImpl struct{}

func NewWaitGroupModule() Module {
	return &waitGroupModuleImpl{}
}

func (wgm *waitGroupModuleImpl) FindAvgFromFile(filename string) error {
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

	n := len(records) / 50

	wg := sync.WaitGroup{}
	sumArr := [50]float64{}
	for i := range 50 {
		wg.Add(1)
		go func(sumArr *[50]float64, err *error) {
			defer wg.Done()
			sum := 0.0
			for _, record := range records[i*n : (i+1)*n] {
				value, parseErr := strconv.ParseFloat(record[0], 64)
				if parseErr != nil {
					*err = parseErr
					return
				}
				sum += value
			}
			(*sumArr)[i] = sum
		}(&sumArr, &err)
	}

	wg.Wait()

	if err != nil {
		return err
	}

	sum := 0.0
	for i := range 50 {
		sum += sumArr[i]
	}

	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("\nWait Group module average: %f", sum/float64(len(records))))
	fmt.Println(fmt.Sprintf("Wait Group module read 1000 rows of CSV file took %s", elapsed))

	return nil
}
