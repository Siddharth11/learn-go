package main

import (
	"fmt"
	"reflect"
)

var name string = "Siddharth"
var namePointer = &name

func main() {
	fmt.Println("Variable \"name\" has value:", name, "and type:", reflect.TypeOf(name), "address", namePointer, "and pointer value", *namePointer)
}
