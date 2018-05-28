package sum

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) []int {
	length := len(numberToSum)
	sums := make([]int, length)

	for key, value := range numberToSum {
		sums[key] = Sum(value)
		fmt.Print("-------------------------\n")
		fmt.Printf("%d \n",key)
	}
	return sums
}
