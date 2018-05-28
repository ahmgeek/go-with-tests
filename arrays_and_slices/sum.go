package sum

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numberToSum ...[]int) (sums []int) {
	for _, numbers := range numberToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
