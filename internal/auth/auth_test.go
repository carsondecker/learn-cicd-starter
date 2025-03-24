package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		headerKey  string
		headerVal  string
		givesError bool
		output     string
		errorMsg   string
	}

	cases := []test{
		{
			headerKey:  "Not-Auth",
			headerVal:  "ApiKey asdfasdf",
			givesError: true,
			output:     "",
			errorMsg:   "no authorization header included",
		},
		{
			headerKey:  "Authorization",
			headerVal:  "Bearer asdfasdf",
			givesError: true,
			output:     "",
			errorMsg:   "malformed authorization header",
		},
		{
			headerKey:  "Authorization",
			headerVal:  "ApiKey:asdfasdf",
			givesError: true,
			output:     "",
			errorMsg:   "malformed authorization header",
		},
		{
			headerKey:  "Authorization",
			headerVal:  "ApiKey asdfasdf",
			givesError: false,
			output:     "asdfasdf",
			errorMsg:   "",
		},
	}

	for i, c := range cases {
		req, _ := http.NewRequest("GET", "http://example.com", nil)
		req.Header.Add(c.headerKey, c.headerVal)

		apiKey, err := GetAPIKey(req.Header)
		if err != nil {
			if !c.givesError {
				t.Errorf("case %d returned unexpected error %v", i, err)
			}
			if c.errorMsg != err.Error() {
				t.Errorf("case %d expected error message '%s' but got '%s'", i, c.errorMsg, err.Error())
			}
		}

		if !c.givesError {
			if apiKey != c.output {
				t.Errorf("case %d expected output '%s' but got '%s'", i, c.output, apiKey)
			}
		}
	}
}
