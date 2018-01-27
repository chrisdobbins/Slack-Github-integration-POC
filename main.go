package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/circlewebhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success!")
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("bad request")
			return
		}
		fmt.Println(fmt.Sprintf("%+v", body))

	})
	http.ListenAndServe(":8080", nil)
}
