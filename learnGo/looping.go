package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		for j := i + 1; j < 10; j++ {
			fmt.Print(j)
		}
	}
}
