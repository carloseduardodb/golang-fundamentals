package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = *template.Must(template.ParseGlob("*.html"))

type user struct {
	Name  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := user{
		Name:  "Tom",
		Email: "tom@hotmail.com",
	}
	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {
	http.HandleFunc("/home", home)

	fmt.Println("Server started on: http://localhost:3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
}
