package main

import "fmt"

const (
	firstName, lastName = "John", "Doe"
	age                 = 32
)

var (
	address    string
	postalCode int
)

func main() {
	fmt.Println(firstName, lastName, age)
	address = "qwerty Avenue"
	postalCode = 12345
	fmt.Println(address, postalCode)
}
