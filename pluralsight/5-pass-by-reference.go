package main

import "fmt"

func main() {
	var name string = "Siddharth"

	fmt.Println("\nValue of \"name\" before update func:", name)

	updateName(&name)

	fmt.Println("\nValue of \"name\" after update func:", name)
}

func updateName(name *string) string {
	*name = "Muskan"

	fmt.Println("\nValue of \"name\" inside update func:", *name)

	return *name
}
