package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CircleResp struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	VCS_URL  string `json:"vcs_url"`
	BuildURL string `json:"build_url"`
	BuildNum string `json:"build_num"`
	Branch   string `json:"branch"`
}

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
		notification := CircleResp{}
		fmt.Println(fmt.Sprintf("%+v", body))
		err = json.Unmarshal(body, &notification)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(notification.Payload)

	})
	http.ListenAndServe(":8080", nil)
}
