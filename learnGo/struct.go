package main

import "fmt"

type man struct {
	name string
	age  int
	job  string
}

func main() {
	Lan := man{
		name: "Lan",
		age:  22,
		job:  "Hr",
	}
	fmt.Print("Name :" + Lan.name)
}
