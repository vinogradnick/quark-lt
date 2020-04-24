package main

import (
	"fmt"
	"net/http"
)
var counter=0
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Println( counter)
	})
	http.ListenAndServe(":1488", nil)
}