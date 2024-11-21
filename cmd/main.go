package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Porsche-ths/simple-goroutine/internal/modules"
)

func readInput(reader *bufio.Reader) int {
	fmt.Print("Enter Jobs Number: ")
	inputText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error reading input please try again")
		return readInput(reader)
	}

	input := strings.Replace(inputText, "\n", "", -1)
	jobsNum, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Print("The input must be an integer\n")
		return readInput(reader)
	}

	return int(jobsNum)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	jobsNum := readInput(reader)

	simpleModule := modules.NewSimpleModule()
	simpleErr := simpleModule.FindAvgFromFile("./files/data.csv", jobsNum)
	if simpleErr != nil {
		log.Fatalf("Error: %s", simpleErr)
	}

	waitGroupModule := modules.NewWaitGroupModule()
	waitGroupErr := waitGroupModule.FindAvgFromFile("./files/data.csv", jobsNum)
	if waitGroupErr != nil {
		log.Fatalf("Error: %s", waitGroupErr)
	}

	channelModule := modules.NewChannelModule()
	channelErr := channelModule.FindAvgFromFile("./files/data.csv", jobsNum)
	if channelErr != nil {
		log.Fatalf("Error: %s", channelErr)
	}

	workerPoolModule := modules.NewWorkerPoolModule()
	workerPoolErr := workerPoolModule.FindAvgFromFile("./files/data.csv", jobsNum)
	if workerPoolErr != nil {
		log.Fatalf("Error: %s", workerPoolErr)
	}
}
