// Client/Server AJAX JSON Communication using golang web-server and JQuery
// Visit: http://127.0.0.1:8080
package main

import (
    "fmt"
    "log"
    "encoding/json"
    "html/template"
    "net/http"
    // "path"
)

type Data struct {
    UserInput string
    ElizaOutput string
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/eliza", 301)
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    // fp := path.Join("templates", "ajax-json.html")
    tmpl, err := template.ParseFiles("eliza.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    //parse request to struct
    var d Data
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    fmt.Println(d.UserInput)

    d.ElizaOutput = d.UserInput
    // create json response from struct
    a, err := json.Marshal(d)

    fmt.Println(d.ElizaOutput)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.Write(a)
}

func main() {
    http.HandleFunc("/", redirect)
    http.HandleFunc("/eliza", defaultHandler)
    http.HandleFunc("/ajax", ajaxHandler)

    fmt.Println("Server running at port 8080...")

    err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

