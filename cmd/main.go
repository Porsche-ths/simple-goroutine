package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Porsche-ths/simple-goroutine/internal/modules"
)

func main() {
	simpleStart := time.Now()

	simpleModule := modules.NewSimpleModule()
	simpleAvg, err := simpleModule.FindAvgFromfile("./files/data.csv")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	simpleElapsed := time.Since(simpleStart)

	fmt.Println(fmt.Sprintf("\nSimple module average: %f", simpleAvg))
	fmt.Println(fmt.Sprintf("Read 1000 rows of CSV file took %s", simpleElapsed))
}
