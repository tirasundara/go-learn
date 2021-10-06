package slices

// Sum returns sum of any numbers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
