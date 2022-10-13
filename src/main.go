package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

//ToDo: Change ProblemDetails
//ToDo:
//- [ ] yaml fertig machen
//- [ ] error/success messages
//- [ ] controller.go schreiben, interface für POST anlegen
//- [ ] Service webresponder fuer go(net/http?) suchen und injecten -> Martin fragen
//- [ ] SuccessMessage fehlt in dtos?
//- [ ] 1 function in interface: auf controllerschicht
//- [ ] prüft request
//- [ ] schreibt daten in Go-object (mit dto)
//- [ ] benutzt go-object um antwort zu generieren
//- [ ] module.go (oder router in main definieren?)
