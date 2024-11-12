package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	elapsed := time.Since(start)
	log.Printf("Read 1000 rows of CSV file took %s", elapsed)
}