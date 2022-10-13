package main

import (
	"encoding/json"
	"net/http"

	"github.com/moschmuc/test-project/src/dtos"
)

func helloGreetingHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Hello World!")
	var gr dtos.GreetingRequest
	err := json.NewDecoder(r.Body).Decode(&gr)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func goodbyeHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Goodbye World!")
}
