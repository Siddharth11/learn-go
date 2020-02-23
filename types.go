package main

import (
	"fmt"
	"reflect"
)

var name string = "Siddharth"
var age int = 25

func main() {
	name2 := "Muskan"
	age2 := 22

	fmt.Println("Variable \"name\" has value:", name, "and type:", reflect.TypeOf(name))

	fmt.Println("\nVariable \"age\" has value:", age, "and type:", reflect.TypeOf(age))

	fmt.Println("\nVariable \"name2\" has value:", name2, "and inferred type:", reflect.TypeOf(name2))

	fmt.Println("\nVariable \"age2\" has value:", age2, "and type:", reflect.TypeOf(age2))
}
