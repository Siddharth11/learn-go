package main

import (
	"fmt"
	"strings"
)

func main() {
	const name1 string = "Siddharth"
	const name2 string = "Muskan"

	fmt.Println(converter(name1, name2))
}

func converter(str1 string, str2 string) (string, string) {
	return strings.Title(str1), strings.ToUpper(str2)
}
