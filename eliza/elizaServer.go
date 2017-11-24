package eliza

import (
	"encoding/json"
	"html/template"
	"net/http"
	"log"
)

// Structs for using JSON
type UserInput struct {
	UserMessage string
}

type ServerOutput struct {
	ServerMessage string
}

// Redirect to '/eliza'
func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:8080/eliza", 301)
}

// Default Request Handler
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	// fp := path.Join("templates", "ajax-json.html")
	tmpl, err := template.ParseFiles("eliza.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

} // defaultHandler

// AJAX Request Handler
func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	//parse request to struct
	var userInput UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// log.Println("DEBUG: eliza.AjaxHandler: userInput.UserMessage: " + userInput.UserMessage)

	// Create a new instance of Eliza.
	eliza := FromFiles("data/responses.txt", "data/substitutions.txt")

	var elizaOutput ServerOutput
	elizaOutput.ServerMessage = eliza.RespondTo(userInput.UserMessage)

	// log.Println("DEBUG: eliza.AjaxHandler: elizaOutput.ElizaMessage: " + elizaOutput.ElizaMessage)

	// create json response from struct
	reply, err := json.Marshal(elizaOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(reply)

} // AjaxHandler

func Start() {
	http.HandleFunc("/ajax", AjaxHandler)
	http.Handle("/", http.FileServer(http.Dir("./html")))

	log.Printf("Starting server at port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}