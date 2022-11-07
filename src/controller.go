package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/moschmuc/test-project/src/dtos"
)

//var em dtos.ErrorMessage

func helloGreetingHandler(w http.ResponseWriter, r *http.Request) {
	var gr dtos.GreetingRequest
	err := json.NewDecoder(r.Body).Decode(&gr)
	GreetingRequestHandler(w, err, gr)

}

func GreetingRequestHandler(w http.ResponseWriter, err error, gr dtos.GreetingRequest) int {

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return isValid(w, gr)
}

func isValid(w http.ResponseWriter, gr dtos.GreetingRequest, Errs url.Values) int {
	var responseCode int
	if (*gr.Salutation == "" || *gr.Salutation == "Divers") && (*gr.FirstName == "" || gr.LastName == "") {
		responseCode = badMove(w, responseCode)
		Errs.Add("salutation", "the salutation ist required")
	} else if (*gr.FirstName == "" && *gr.Salutation == "") || (*gr.FirstName == "" && gr.LastName == "") {
		responseCode = badMove(w, responseCode)
	} else if *gr.FirstName != "" && gr.LastName == "" {
		responseCode = badMove(w, responseCode)
	} else if *gr.FirstName != "" && gr.LastName != "" {
		fmt.Println("Hello", *gr.FirstName, gr.LastName)
		responseCode = goodMove(w, responseCode)
	} else if *gr.Salutation != "" && gr.LastName != "" {
		fmt.Println("Hello", *gr.Salutation, gr.LastName)
		responseCode = goodMove(w, responseCode)
	}
	return responseCode
}

func goodMove(w http.ResponseWriter, responseCode int) int {
	responseCode = http.StatusOK
	successMessage := dtos.SuccessMessage{Message: "looks good dude"}
	w.WriteHeader(responseCode)
	fmt.Println(successMessage)
	return responseCode
}

func badMove(w http.ResponseWriter, responseCode int) int {
	responseCode = http.StatusBadRequest
	ErrorMessage := dtos.ErrorMessage{Error: "Please enter at least your last name and your first name or your salutation"}
	w.WriteHeader(responseCode)
	fmt.Println(ErrorMessage)
	return responseCode
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
