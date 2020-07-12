package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./this-file-does-not-exist.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}
}