package sum

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberToSum ...[]int) (sums []int) {
	for _, numbers := range numberToSum {
		fmt.Print("-------------------------\n")
		fmt.Printf("%d \n",numbers)
		
		sums = append(sums, Sum(numbers))
	}
	return sums
}
