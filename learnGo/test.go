// server.go
/*
package main

import (
	"fmt"
	_ "encoding/json"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Welcome to our website")
	return
}

func main(){
	http.HandleFunc("/", welcome)

	http.ListenAndServe(":8080", nil)
}
*/
package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x)
}
