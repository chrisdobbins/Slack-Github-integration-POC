package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CircleResp struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	VCS_URL  string `json:"vcs_url"`
	BuildURL string `json:"build_url"`
	BuildNum int    `json:"build_num"`
	Branch   string `json:"branch"`
	Outcome  string `json:"outcome"`
}

type GHEvent struct {
	Release Release `json:"release"`
}

type Release struct {
	TagName string `json:"tag_name"`
}

func main() {
	appUrl := os.Getenv("SLACK_URL")

	http.HandleFunc("/githubwebhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success!")
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("bad request")
			return
		}
		ghEvt := GHEvent{}
		err = json.Unmarshal(body, &ghEvt)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(fmt.Sprintf("%+v", ghEvt))
		http.Post(appUrl, "application/json", bytes.NewBuffer([]byte(`{"text": "testing, testing..."}`)))
	})
	http.HandleFunc("/circlewebhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success!")
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("bad request")
			return
		}
		notification := CircleResp{}
		err = json.Unmarshal(body, &notification)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(notification.Payload)

	})
	http.ListenAndServe(":8080", nil)
}
