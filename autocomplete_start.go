package main

import (
	"log"
	"net/http"
)

// program entry point
func main() {

	local := true // TODO: implement prod version
	url := ""
	port := ":9000"

	if local == true {
		url = "localhost"
	}

	// set handler
	http.HandleFunc("/autocomplete", autocompleteHandler)

	// start http listener
	log.Println("Autocomplete service listening at: " + url + port + "/autocomplete")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
