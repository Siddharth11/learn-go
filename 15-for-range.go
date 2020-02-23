package main

import (
	"fmt"
)

func main() {

	names := []string{"Ram", "Shayam", "Ghanshayam"}

	for _, personName := range names {
		fmt.Println(personName)
	}
}