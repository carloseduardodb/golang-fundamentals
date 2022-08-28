package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body><h1>Hello World</h1></body></html>`))
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body><h1>User</h1></body></html>`))
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/user", user)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
