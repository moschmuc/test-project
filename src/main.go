package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/v1/greeting", helloGreetingHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
