package modules

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type optimizedModuleImpl struct{}

func NewOptimizedModule() Module {
	return &optimizedModuleImpl{}
}

func readAndSum(reader *csv.Reader, returnErr *error) float64 {
	sum := 0.0
	for {
		row, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				*returnErr = err
			}
			break
		} else if len(row) == 0 {
			break
		}

		value, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			*returnErr = err
			break
		}
		sum += value
	}
	return sum
}

func worker(jobsNum int, result chan<- float64, reader *csv.Reader, returnErr *error) {
	for i := 0; i < jobsNum; i++ {
		result <- readAndSum(reader, returnErr)
	}
}

func (om *optimizedModuleImpl) FindAvgFromFile(filename string, jobsNum int) error {
	start := time.Now()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	jobs := make(chan int, jobsNum)
	result := make(chan float64, jobsNum)

	go worker(jobsNum, result, reader, &err)
	if err != nil {
		return err
	}

	for i := 0; i < jobsNum; i++ {
		jobs <- i
	}
	close(jobs)

	sum := 0.0
	for i := 0; i < jobsNum; i++ {
		sum += <-result
	}
	close(result)

	elapsed := time.Since(start)

	fmt.Println(fmt.Sprintf("\nOptimized module average: %f", sum/50000000.0))
	fmt.Println(fmt.Sprintf("Optimized module read 50M rows of CSV file took %s", elapsed))

	return nil
}
