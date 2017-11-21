package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

type elizaData struct {
	UserInput string
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/eliza", 301)
}

func handlerEliza(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "eliza.html")
	// data := &elizaData{UserInput: "some user input"}
	// templ, _ := template.ParseFiles("eliza.html")

	user := r.FormValue("user-input")

	data := elizaData{UserInput: "some input"}

	fp := path.Join("templates", "form.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Pass template to http.ResponseWriter.
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// templ.Execute(w, data)
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
