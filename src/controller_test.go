package main

import (
	"testing"

	"github.com/moschmuc/test-project/src/dtos"
)

func Test_validateRequest(t *testing.T) {
	type (
		testCase struct {
			name    string
			gr      dtos.GreetingRequest
			wantErr bool
		}
		testCases []testCase
	)
	tests := testCases{
		{ //if 1, Z. 64
			name: "NOK:first name salutation",
			gr: dtos.GreetingRequest{
				FirstName:  stringToNilableString("Momme"),
				LastName:   "",
				Salutation: stringToNilableGreetingRequestSalutation("Herr"),
			},
			wantErr: true,
		},
		{
			name: "OK:last name salutation",
			gr: dtos.GreetingRequest{
				FirstName:  nil,
				LastName:   "Petersen",
				Salutation: stringToNilableGreetingRequestSalutation("Frau"),
			},
			wantErr: false,
		},
		//{
		//	name: "NOK: first name",
		//	gr: dtos.GreetingRequest{
		//		FirstName:  stringToNilableString("Peter"),
		//		LastName:   "",
		//		Salutation: nil,
		//	},
		//	wantErr: true,
		//},
		{
			name: "NOK: salutation Divers first name",
			gr: dtos.GreetingRequest{
				FirstName:  stringToNilableString("Michael"),
				LastName:   "",
				Salutation: stringToNilableGreetingRequestSalutation("Divers"),
			},
			wantErr: true,
		},
		{
			name: "NOK: salutation Divers last name",
			gr: dtos.GreetingRequest{
				FirstName:  nil,
				LastName:   "Mueller",
				Salutation: stringToNilableGreetingRequestSalutation("Divers"),
			},
			wantErr: true,
		}, // else if 2, Z. 67
		//{
		//	name: "NOK: all empty",
		//	gr: dtos.GreetingRequest{
		//		FirstName:  stringToNilableString(""),
		//		LastName:   "",
		//		Salutation: nil,
		//	},
		//	wantErr: true,
		//},
		//{
		//	name: "OK: first name last name",
		//	gr: dtos.GreetingRequest{
		//		FirstName:  stringToNilableString("Jenny"),
		//		LastName:   "Wurst",
		//		Salutation: nil,
		//	},
		//	wantErr: false,
		//},
		//{
		//	name: "NOK: last name",
		//	gr: dtos.GreetingRequest{
		//		FirstName:  stringToNilableString(""),
		//		LastName:   "Wurst",
		//		Salutation: nil,
		//	},
		//	wantErr: true,
		//},
		{
			name: "NOK: wrong salutation",
			gr: dtos.GreetingRequest{
				FirstName:  stringToNilableString(""),
				LastName:   "",
				Salutation: stringToNilableGreetingRequestSalutation("Falsche Begruessung"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateRequest(tt.gr); (err != nil) != tt.wantErr {
				t.Errorf("validateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
