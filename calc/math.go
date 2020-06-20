package calc

import "errors"

// return sum of two integer
func Add(numbers ...int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("invalid number of args provided")
	}
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return sum, nil
}
