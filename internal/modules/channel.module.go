package modules

import (
	"encoding/csv"
	"fmt"
	"os"
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

	sumArr := make(chan float64, jobsNum)
	for i := range jobsNum {
		go func() {
			sumArr <- calculateSum(records[i*n:(i+1)*n], &err)
		}()
	}

	if err != nil {
		return err
	}

	sum := 0.0
	for i := 0; i < jobsNum; i++ {
		sum += <-sumArr
	}

	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("\nChannel module average: %f", sum/float64(len(records))))
	fmt.Println(fmt.Sprintf("Channel module read 10M rows of CSV file took %s", elapsed))

	return nil
}
