package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	firstName := "John"
	updateName(firstName)
	fmt.Println(firstName)

	firstName2 := "Mike"
	updateNamePointer(&firstName2)
	fmt.Println(firstName2)
}

func calc(number1 string, number2 string) (int, int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum := int1 + int2
	mul := int1 * int2
	return sum, mul
}

func updateName(name string) {
	name = "David"
}

func updateNamePointer(name *string) {
	*name = "David"
}
