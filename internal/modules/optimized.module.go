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

func readAndSum(reader *csv.Reader, rowCount *float64, returnErr *error) float64 {
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
		*rowCount++
		sum += value
	}

	return sum
}

func optimizedWorker(jobsNum int, result chan<- float64, reader *csv.Reader, rowCount *float64, returnErr *error) {
	for i := 0; i < jobsNum; i++ {
		result <- readAndSum(reader, rowCount, returnErr)
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

	rowCount := 0.0
	jobs := make(chan int, jobsNum)
	result := make(chan float64, jobsNum)

	go optimizedWorker(jobsNum, result, reader, &rowCount, &err)
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

	fmt.Println(fmt.Sprintf("\nOptimized module average: %f", sum/rowCount))
	fmt.Println(fmt.Sprintf("Optimized module read 50M rows of CSV file took %s", elapsed))

	return nil
}
