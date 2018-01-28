package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type GHEvent struct {
	Release Release `json:"release"`
}

type Release struct {
	TagName string `json:"tag_name"`
}

func main() {
	appUrl := os.Getenv("SLACK_URL")
	if appUrl == "" {
		log.Fatal("SLACK_URL not set")
	}

	http.HandleFunc("/githubwebhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success!")
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
		ghEvt := GHEvent{}
		err = json.Unmarshal(body, &ghEvt)
		if err != nil {
			log.Println(err.Error())
			return
		}
		tag := ghEvt.Release.TagName
		fmt.Println(fmt.Sprintf("%+v", ghEvt))
		http.Post(appUrl, "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"text": "New version number is: %s"}`, tag))))
	})

	http.ListenAndServe(":8080", nil)
}
