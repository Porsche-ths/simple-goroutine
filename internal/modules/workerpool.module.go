package modules

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type workerPoolModuleImpl struct{}

func NewWorkerPoolModule() Module {
	return &workerPoolModuleImpl{}
}

func (wpm *workerPoolModuleImpl) FindAvgFromFile(filename string, jobsNum int) error {
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

	jobs := make(chan int, jobsNum)
	sumArr := make(chan float64, jobsNum)
	for i := range jobsNum {
		go func(sumArr *chan float64, err *error) {
			sum := 0.0
			for _, record := range records[i*n : (i+1)*n] {
				value, parseErr := strconv.ParseFloat(record[0], 64)
				if parseErr != nil {
					*err = parseErr
					return
				}
				sum += value
			}
			(*sumArr) <- sum
		}(&sumArr, &err)
	}

	if err != nil {
		return err
	}

	for i := range jobsNum {
		jobs <- i
	}
	close(jobs)

	sum := 0.0
	for i := 0; i < jobsNum; i++ {
		sum += <-sumArr
	}

	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("\nWorker Pool module average: %f", sum/float64(len(records))))
	fmt.Println(fmt.Sprintf("Worker Pool module read 1000 rows of CSV file took %s", elapsed))

	return nil
}
