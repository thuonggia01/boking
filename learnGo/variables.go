package main

import "fmt"

func main() {
	var a int = 50
	b := 50
	fmt.Println(a)
	fmt.Println(b)
	var c bool
	c = false
	fmt.Printf("%v,%T\n", c, c)
	var (
		name    string = "aomacalada"
		age     int    = 100
		address string = "bon phuong"
	)
	fmt.Printf("%v,%T \n", name, name)
	fmt.Printf("%v,%T\n", age, age)
	fmt.Printf("%v,%T\n", address, address)
}
