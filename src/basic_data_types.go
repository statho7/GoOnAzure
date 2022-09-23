package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	e        = 2.71828
	Avogadro = 6.02214129e23
	Planck   = 6.62606957e-34
)

var (
	integer8  int8  = 127
	integer16 int16 = 32767
	integer32 int32 = 2147483647
	integer64 int64 = 9223372036854775807

	defaultInt     int
	defaultFloat32 float32
	defaultFloat64 float64
	defaultBool    bool
	defaultString  string
)

func previous_main() {
	fmt.Println(integer8, integer16, integer32, integer64)

	//.\main.go:19:14: invalid operation: integer16 + integer32 (mismatched types int16 and int32)
	//fmt.Println(integer16 + integer32)

	rune := 'G'
	fmt.Println(rune)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	fmt.Println(e, Avogadro, Planck)

	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)

	fmt.Println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
