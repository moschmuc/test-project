package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moschmuc/test-project/src/dtos"
)

//var em dtos.ErrorMessage

func helloGreetingHandler(w http.ResponseWriter, r *http.Request) {
	//write(w, "Hello Jenny!")
	var gr dtos.GreetingRequest
	err := json.NewDecoder(r.Body).Decode(&gr)
	fmt.Println(*gr.Salutation, *gr.FirstName, gr.LastName)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Goodbye World!")
}

//func (*Controller) unmarshal(r *web.Request, value any) error {
//	if r.Request() == nil {
//		return errMissingHTTPRequest
//	}
//
//	if r.Request().Body == nil || r.Request().Body == http.NoBody {
//		return errEmptyBody
//	}
//
//	err := json.NewDecoder(r.Request().Body).Decode(&Greeting)
//	if err != nil {
//		return fmt.Errorf("request body could not be unmarshalled: %w", err)
//	}
//	return nil
//}
