package main

import "fmt"

func main() {
	num1 := 25
	num2 := 5
	remainder:= num1 % num2

	switch remainder {
		case 0:
			fmt.Println("\nnum1 is divisible by num2")
		default:
			fmt.Println("\nnum1 is not divisible by num2")
	}
}