package Arrays_and_slices

func SumAll(numbersToSums ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSums {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
