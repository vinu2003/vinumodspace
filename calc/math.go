package calc

// return sum of two integer
func Add(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	return sum
}
