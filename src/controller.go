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
	err = validateRequest(gr)
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

	w.WriteHeader(http.StatusOK)
	greetingString := createGreeting(gr)

	err = json.NewEncoder(w).Encode(
		dtos.SuccessMessage{
			Message: greetingString,
		},
	)
	return
}

func createGreeting(gr dtos.GreetingRequest) string {
	var greetingString string

	if *gr.FirstName != "" && gr.LastName != "" {
		greetingString = fmt.Sprintf("Hello %s %s", *gr.FirstName, gr.LastName)
	} else {
		greetingString = fmt.Sprintf("Hello %s %s", *gr.Salutation, gr.LastName)
	}
	return greetingString
}

func validateRequest(gr dtos.GreetingRequest) error {
	if gr.LastName != "" && isEmpty(gr.FirstName) &&
		(grIsEmpty(gr.Salutation) || *gr.Salutation == dtos.Divers) {
		return errors.New("please enter at least a salutation (Frau/Herr) and a last name or a first name and a last name")

	} else if !grIsEmpty(gr.Salutation) && (gr.FirstName != nil && *gr.FirstName == "" && gr.LastName == "") {
		return errors.New("first and last name are missing")

	} else if !isEmpty(gr.FirstName) && gr.LastName == "" {
		return errors.New("last name is missing")

	} else if !grIsEmpty(gr.Salutation) && *gr.Salutation != dtos.Frau && *gr.Salutation != dtos.Herr && *gr.Salutation != dtos.Divers {
		return fmt.Errorf("%s is not a valid salutation", *gr.Salutation)

	} else if *gr.FirstName == "" && gr.LastName == "" && *gr.Salutation == "" {
		return errors.New("none of the attributes are filled")
	}
	return nil
}

func isEmpty[T string | dtos.GreetingRequestSalutation](t *T) bool {
	return t == nil || *t == ""
}

func grIsEmpty(s *dtos.GreetingRequestSalutation) bool {
	return s == nil || *s == ""
}
func stringToNilableString(value string) *string {
	var defaultValue string
	if value == defaultValue {
		return nil
	}

	result := value
	return &result
}
func stringToNilableGreetingRequestSalutation(value string) *dtos.GreetingRequestSalutation {
	var defaultValue string
	if value == defaultValue {
		return nil
	}

	result := dtos.GreetingRequestSalutation(value)
	return &result
}
