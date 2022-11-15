package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/moschmuc/test-project/src/dtos"
)

func GreetingRequestHandler(w http.ResponseWriter, r *http.Request) {
	var gr dtos.GreetingRequest

	err := json.NewDecoder(r.Body).Decode(&gr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = checkErrors(gr)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		err = json.NewEncoder(w).Encode(
			dtos.ErrorMessage{
				Error: err.Error(),
			},
		)
		if err != nil {
			panic("marshalling error response failed, something is really wrong here")
		}
		return
	}
	//ToDo: implement func Sprintf (done)
	//ToDo: set variable var bla string (done)
	//ToDo: export if cases (See func checkErrors)

	w.WriteHeader(http.StatusOK)
	greetingString := checkSuccess(gr)

	err = json.NewEncoder(w).Encode(
		dtos.SuccessMessage{
			Message: greetingString,
		},
	)
	return
}

func checkSuccess(gr dtos.GreetingRequest) string {
	var greetingString string

	if *gr.FirstName != "" && gr.LastName != "" {
		greetingString = fmt.Sprintf("Hello %s %s", *gr.FirstName, gr.LastName)
	} else {
		greetingString = fmt.Sprintf("Hello %s %s", *gr.Salutation, gr.LastName)
		// *gr.Salutation != "" && gr.LastName != "" {
	}
	return greetingString
}

// ToDo: Start here with unit tests
func checkErrors(gr dtos.GreetingRequest) error {
	if (*gr.Salutation == "" || *gr.Salutation == dtos.Divers) && (*gr.FirstName == "" || gr.LastName == "") {
		err := errors.New("please enter at least a salutation (Frau/Herr) and a last name or a first name and a last name")
		return err

	} else if *gr.FirstName == "" && gr.LastName == "" {
		err := errors.New("first and last name are missing")
		return err

	} else if *gr.FirstName != "" && gr.LastName == "" {
		err := errors.New("last name is missing")
		return err

	} else if *gr.Salutation != dtos.Frau && *gr.Salutation != dtos.Herr && *gr.Salutation != dtos.Divers && *gr.Salutation != "" {
		err := fmt.Errorf("%s is not a valid salutation", *gr.Salutation)
		return err
	}

	return nil
}
