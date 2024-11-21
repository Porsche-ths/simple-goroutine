package main

import (
	"log"

	"github.com/Porsche-ths/simple-goroutine/internal/modules"
)

func main() {
	simpleModule := modules.NewSimpleModule()
	simpleErr := simpleModule.FindAvgFromFile("./files/data.csv")
	if simpleErr != nil {
		log.Fatalf("Error: %s", simpleErr)
	}

	waitGroupModule := modules.NewWaitGroupModule()
	waitGroupErr := waitGroupModule.FindAvgFromFile("./files/data.csv")
	if waitGroupErr != nil {
		log.Fatalf("Error: %s", waitGroupErr)
	}

	channelModule := modules.NewChannelModule()
	channelErr := channelModule.FindAvgFromFile("./files/data.csv")
	if channelErr != nil {
		log.Fatalf("Error: %s", channelErr)
	}
}
