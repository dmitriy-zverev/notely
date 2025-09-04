package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header1 := http.Header{}
	header1.Set("Authorization", "ApiKey some-auth-key")

	header2 := http.Header{}

	cases := []http.Header{
		header1,
		header2,
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			apiKey, err := GetAPIKey(c)
			if err != nil &&
				!(fmt.Sprintf("%v", err) == "no authorization header included" ||
					fmt.Sprintf("%v", err) == "malformed authorization header") {
				t.Errorf("unexpected error")
				t.Fail()
			}

			if i == 0 && apiKey != "some-auth-key" {
				t.Error(c)
				t.Errorf("cannot get correct auth key from header")
				t.Fail()
			}
		},
		)
	}
}
