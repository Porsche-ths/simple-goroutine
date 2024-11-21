package modules

type Module interface {
	FindAvgFromFile(filename string) error
}