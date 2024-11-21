package modules

import "strconv"

type Module interface {
	FindAvgFromFile(filename string, jobsNum int) error
}

func calculateSum(records [][]string, err *error) float64 {
	sum := 0.0
	for _, record := range records {
		value, parseErr := strconv.ParseFloat(record[0], 64)
		if parseErr != nil {
			*err = parseErr
			return 0
		}
		sum += value
	}
	return sum
}
