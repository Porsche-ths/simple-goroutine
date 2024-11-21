package modules

type Module interface {
	FindAvgFromFile(filename string, jobsNum int) error
}