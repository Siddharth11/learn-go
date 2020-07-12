package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxInt(1, -1, 2, 4, 4, 99, 199, 200))
}

func getMaxInt(numbers ...int) int {
	var maxValue int = numbers[0]

	for _, num := range numbers {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}
