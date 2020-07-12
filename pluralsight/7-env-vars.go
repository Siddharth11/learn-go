package main

import (
	"fmt"
	"os"
)

func main() {
	var osUser string = os.Getenv("USER")

	fmt.Println("My OS username is:", osUser)
}
