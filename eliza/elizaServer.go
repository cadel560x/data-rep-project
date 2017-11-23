package eliza

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// HTTP server section

// Structs for using JSON
type UserInput struct {
	UserMessage string
}

type ServerOutput struct {
	ServerMessage string
}

// Redirect to '/eliza'
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

} // defaultHandler

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	//parse request to struct
	var userInput UserInput
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// fmt.Println("DEBUG: ajaxHandler: userInput.UserMessage: " + userInput.UserMessage)

	// Create a new instance of Eliza.
	eliza := FromFiles("data/responses.txt", "data/substitutions.txt")

	var elizaOutput ServerOutput
	// elizaOutput.ElizaMessage = ElizaResponse(userInput.UserMessage)
	elizaOutput.ServerMessage = eliza.RespondTo(userInput.UserMessage)

	// fmt.Println("DEBUG: ajaxHandler: elizaOutput.ElizaMessage: " + elizaOutput.ElizaMessage)

	// create json response from struct
	reply, err := json.Marshal(elizaOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(reply)

} // ajaxHandler
