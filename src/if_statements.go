package main

import "fmt"

func somenumber() int {
	return -7
}
func main() {
	num := somenumber()
	if num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println(num)
}
