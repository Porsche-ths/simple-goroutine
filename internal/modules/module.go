package modules

type Module interface {
	FindAvgFromfile(filename string) (float64, error)
}