package main

import "fmt"

func main() {
    num1 := 25
    num2 := 5

    if remainder:= num1 % num2; remainder == 0 {
        fmt.Println("\nnum1 is divisible by num2")
    } else {
        fmt.Println("\nnum1 is not divisible by num2")
    }
}