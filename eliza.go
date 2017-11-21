package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/eliza", 301)
}

func handlerEliza(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "eliza.html")

	templ, _ := template.ParseFiles("eliza.html")

	templ.Execute(w, interface{})
	// templ.Execute(w, placeholders)
}

func main() {
	http.HandleFunc("/", redirect)
	http.HandleFunc("/eliza", handlerEliza)

	fmt.Println("Server running at port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
