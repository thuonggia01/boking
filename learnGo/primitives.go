package main

import (
	"fmt"
	"math"
)

func main() {
	i := 10 == 10
	k := 5 == 10
	fmt.Printf("%v,%T\n", i, i)
	fmt.Printf("%v,%T\n", k, k)

	a := 9
	b := math.Pow(3, 4)
	fmt.Println(a << 3)
	fmt.Println(b)
	c := math.Sqrt(9.3)
	fmt.Println(c)
}
