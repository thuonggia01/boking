package main

import "fmt"

func main() {
	listBienSo := make(map[int]string)
	listBienSo = map[int]string{
		29: "HN",
		30: "HN",
		36: "TH",
		88: "LA",
	}
	fmt.Println(listBienSo)
	listBienSo[99] = "HCM"
	fmt.Println(listBienSo)
	listBienSo[99] = "HCMMMM"
	fmt.Println(listBienSo)
}
