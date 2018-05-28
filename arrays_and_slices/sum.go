package sum

// Sum array numbers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// AllTails : sums all tails in an array.
func AllTails(numberToSum ...[]int) (sums []int) {
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else{

			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
