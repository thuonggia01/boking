package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Sprintln("Type convert")
	var a int = 88
	fmt.Printf("%v,%T\n", a, a)

	var b float64 = float64(a)
	fmt.Printf("%v,%T\n", b, b)

	var c int = int(b)
	fmt.Printf("%v,%T\n", c, c)

	var d string = string(c)
	fmt.Printf("%v,%T\n", d, d)

	k, _ := strconv.Atoi("-55")
	fmt.Printf("%v,%T\n", k, k)

	s := strconv.Itoa(42)
	fmt.Printf("%v,%T\n", s, s)

	// 	b, err := strconv.ParseBool("true")
	// f, err := strconv.ParseFloat("3.1415", 64)
	// i, err := strconv.ParseInt("-42", 10, 64)
	// u, err := strconv.ParseUint("42", 10, 64)
	// s := "2147483647" // biggest int32
	// i64, err := strconv.ParseInt(s, 10, 32)
	// ...
	// i := int32(i64)
	// s := strconv.FormatBool(true)
	// s := strconv.FormatFloat(3.1415, 'E', -1, 64)
	// s := strconv.FormatInt(-42, 16)
	// s := strconv.FormatUint(42, 16)
	// q := strconv.Quote("Hello, 世界")
	// q := strconv.QuoteToASCII("Hello, 世界")
}
