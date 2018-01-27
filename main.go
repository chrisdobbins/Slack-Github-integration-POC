package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/circlewebhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success!")
	})
	http.ListenAndServe(":8080", nil)
}
