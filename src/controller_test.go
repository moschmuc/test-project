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
		{
			name: "NOK:first name salutation",
			gr: dtos.GreetingRequest{
				FirstName:  stringToNilableString("Momme"),
				LastName:   "",
				Salutation: stringToNilableGreetingRequestSalutation("Herr"),
			},
			wantErr: true,
		},
		{
			name: "NOK:",
			gr: dtos.GreetingRequest{
				FirstName:  stringToNilableString(""),
				LastName:   "",
				Salutation: stringToNilableGreetingRequestSalutation(""),
			},
			wantErr: true,
		},
		{
			name: "NOK:",
			gr: dtos.GreetingRequest{
				FirstName:  stringToNilableString(""),
				LastName:   "",
				Salutation: stringToNilableGreetingRequestSalutation(""),
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
