package main

import (
	"encoding/json"
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

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(
			dtos.SuccessMessage{
				Message: "success",
			},
		)
		if *gr.FirstName != "" && gr.LastName != "" {
			fmt.Println("Hello", *gr.FirstName, gr.LastName)

		} else if *gr.Salutation != "" && gr.LastName != "" {
			fmt.Println("Hello", *gr.Salutation, gr.LastName)

		}
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(
			dtos.ErrorMessage{
				Error: err.Error(),
			},
		)
	}
	if err != nil {
		panic("marshalling error response failed, something is really wrong here")
	}
}

func validateRequest(gr dtos.GreetingRequest) error {
	if (*gr.Salutation == "" || *gr.Salutation == dtos.Divers) && (*gr.FirstName == "" || gr.LastName == "") {
		err := fmt.Errorf("please enter at least a salutation (Frau/Herr) and a last name or a first name and a last name")
		fmt.Println(err.Error())
		return err

	} else if *gr.FirstName == "" && gr.LastName == "" {
		err := fmt.Errorf("first and last name are missing")
		fmt.Println(err.Error())
		return err

	} else if *gr.FirstName != "" && gr.LastName == "" {
		err := fmt.Errorf("last name is missing")
		fmt.Println(err.Error())
		return err

	} else if *gr.Salutation != dtos.Frau && *gr.Salutation != dtos.Herr && *gr.Salutation != dtos.Divers {
		err := fmt.Errorf("%s is not a valid salutation", *gr.Salutation)
		fmt.Println(err.Error())
		return err
	}

	return nil
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
