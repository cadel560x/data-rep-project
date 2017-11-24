package main

import (
	// "log"
	// "net/http"

	"./eliza"
)

func main() {
	// Creating a reference to 'eliza.AjaxHandler'
	// var ajaxHandler = eliza.AjaxHandler

	// http.HandleFunc("/ajax", ajaxHandler)
	// http.Handle("/", http.FileServer(http.Dir("./html")))

	// log.Printf("Starting server at port 8080...")
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	eliza.Start();
}
