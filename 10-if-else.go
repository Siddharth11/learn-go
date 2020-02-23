package main

import "fmt"

func main() {
	firstPersonAge := 25
	secondPersonAge := 22

	if firstPersonAge > secondPersonAge {
		fmt.Println("\nFirst person is older than second person")
	} else if secondPersonAge > firstPersonAge {
		fmt.Println("\nSecond person is older than first person")
	} else {
		fmt.Println("\nBoth first and second person are of same age")
	}
}