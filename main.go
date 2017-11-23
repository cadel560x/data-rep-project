package main

import (
	"encoding/json"
	"log"
	"net/http"

	"./eliza"
)

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	// Get instances of the server structs
	var userInput eliza.UserInput
	var serverOutput eliza.ServerOutput

	//parse request to struct
	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// log.Println("DEBUG: ajaxHandler: userInput.UserMessage: " + userInput.UserMessage)

	// Create a new instance of Eliza.
	eliza := eliza.FromFiles("data/responses.txt", "data/substitutions.txt")

	serverOutput.ServerMessage = eliza.RespondTo(userInput.UserMessage)

	// log.Println("DEBUG: ajaxHandler: elizaOutput.ElizaMessage: " + elizaOutput.ElizaMessage)

	// create json response from struct
	reply, err := json.Marshal(serverOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(reply)

} // ajaxHandler

func main() {
	// http.HandleFunc("/eliza", defaultHandler)
	http.HandleFunc("/ajax", ajaxHandler)
	// http.HandleFunc("/css", cssHandler)
	http.Handle("/", http.FileServer(http.Dir("./html")))
	// http.HandleFunc("/", redirect) // Root URN '/' redirects to '/eliza'

	log.Printf("Starting server at port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
