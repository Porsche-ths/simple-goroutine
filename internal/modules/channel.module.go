package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type channelModuleImpl struct{}

func NewChannelModule() Module {
	return &channelModuleImpl{}
}

func (wgm *channelModuleImpl) FindAvgFromFile(filename string, jobsNum int) error {
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

	sumArr := make([]chan float64, jobsNum)
	for i := range jobsNum {
		sumArr[i] = make(chan float64)
		go func(sumArr *[]chan float64, err *error) {
			sum := 0.0
			for _, record := range records[i*n : (i+1)*n] {
				value, parseErr := strconv.ParseFloat(record[0], 64)
				if parseErr != nil {
					*err = parseErr
					return
				}
				sum += value
			}
			(*sumArr)[i] <- sum
		}(&sumArr, &err)
	}

	if err != nil {
		return err
	}

	sum := 0.0
	for i := range jobsNum {
		sum += <-sumArr[i]
	}

	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("\nChannel module average: %f", sum/float64(len(records))))
	fmt.Println(fmt.Sprintf("Channel module read 1000 rows of CSV file took %s", elapsed))

	return nil
}
