package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

type hobby struct {
	Id   int    `json:"id"`
	Item string `json:"item`
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/hello", hello)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, My Google App Engine with Golang")
}

func hello(w http.ResponseWriter, r *http.Request) {
	myHobby := hobby{
		Id:   1,
		Item: "Watch movie",
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(myHobby); err != nil {
		panic(err)
	}
}
