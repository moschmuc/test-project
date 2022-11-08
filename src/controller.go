package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moschmuc/test-project/src/dtos"
)

//var em dtos.ErrorMessage

func GreetingRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	var gr dtos.GreetingRequest
	err = json.NewDecoder(r.Body).Decode(&gr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func handleRequestValidation(w http.ResponseWriter, responseCode int, err error) { //int return value und return responsecode?
	var gr dtos.GreetingRequest
	err = validateRequest(gr)
	if err != nil {
		responseCode = http.StatusBadRequest
		w.WriteHeader(responseCode)
		fmt.Printf(err.Error())

	}
	// successhandling, also goodmove

	if *gr.FirstName != "" && gr.LastName != "" {
		fmt.Println("Hello", *gr.FirstName, gr.LastName)
		returnStatusOK(w, responseCode)

	} else if *gr.Salutation != "" && gr.LastName != "" {
		fmt.Println("Hello", *gr.Salutation, gr.LastName)
		returnStatusOK(w, responseCode)
	}
}
func returnStatusOK(w http.ResponseWriter, responseCode int) int {
	responseCode = http.StatusOK
	//successMessage := dtos.SuccessMessage{Message: "looks good dude"}
	w.WriteHeader(responseCode)
	//fmt.Println(successMessage)
	return responseCode
}

/* func badMove(w http.ResponseWriter, responseCode int) int {
    responseCode = http.StatusBadRequest
    //ErrorMessage := dtos.ErrorMessage{Error: "Please enter at least your last name and your first name or your salutation"}
    w.WriteHeader(responseCode)
    fmt.Println(ErrorMessage)
    return responseCode
} */

func validateRequest(gr dtos.GreetingRequest) error {
	if (*gr.Salutation == "" || *gr.Salutation == "Divers") && (*gr.FirstName == "" || gr.LastName == "") {
		fmt.Errorf("Error: %s, %s or %s are missing", *gr.Salutation, *gr.FirstName, gr.LastName)
	} else if (*gr.FirstName == "" && *gr.Salutation == "") || (*gr.FirstName == "" && gr.LastName == "") {
		fmt.Errorf("testfehler2: %s", gr.Salutation)
	} else if *gr.FirstName != "" && gr.LastName == "" {
		fmt.Errorf("testfehler2: %s", gr.Salutation)

	}
	return nil
}

//statt responsecode bool und *string zurueckgeben (optional)
// Validierungsergebnis-response aus gr-handler anstossen, w nicht notwendig hier

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	write(w, "Goodbye World!")
}

//func (*Controller) unmarshal(r *web.Request, value any) error {
//  if r.Request() == nil {
//      return errMissingHTTPRequest
//  }
//
//  if r.Request().Body == nil || r.Request().Body == http.NoBody {
//      return errEmptyBody
//  }
//
//  err := json.NewDecoder(r.Request().Body).Decode(&Greeting)
//  if err != nil {
//      return fmt.Errorf("request body could not be unmarshalled: %w", err)
//  }
//  return nil
//}
