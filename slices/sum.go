package slices

// Sum returns sum of any numbers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	sums := make([]int, length)

	for i, numbers := range numbersToSum {
		sum := Sum(numbers)
		sums[i] = sum
	}

	return sums
}
