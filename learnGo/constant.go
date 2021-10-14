package main

import "fmt"

func main() {
	const pi = 3.14
	var r float32 = 3
	var p, s float32
	p = pi * 2 * r
	s = pi * r * r

	fmt.Println(pi)
	fmt.Println(s)
	fmt.Println(p)
}
