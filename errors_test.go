package stackongo

import (
	"testing"
)

func TestAllErrors(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/errors", dummyErrorsResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	errors, err := AllErrors(map[string]string{"page": "1"})

	if err != nil {
		t.Error(err.String())
	}

	if len(errors) != 3 {
		t.Error("Number of items wrong.")
	}

	if errors[0].Error_id != 400 {
		t.Error("Error id invalid.")
	}

	if errors[0].Error_name != "bad_parameter" {
		t.Error("error name invalid.")
	}

}

func TestSimulateError(t *testing.T) {
	dummy_server := returnDummyResponseForPath("/2.0/errors/404", dummyErrorResponse, t)
	defer dummy_server.Close()

	//change the host to use the test server
	setHost(dummy_server.URL)

	err := SimulateError(404)

	if err.String() != "no_method: simulated" {
		t.Error("error name invalid.")
	}

}

//Test Data

var dummyErrorResponse string = `
{
  "error_id": 404,
  "error_name": "no_method",
  "error_message": "simulated"
}
`

var dummyErrorsResponse string = `
{
  "items": [
    {
      "error_id": 400,
      "error_name": "bad_parameter",
      "description": "An malformed parameter was passed"
    },
    {
      "error_id": 401,
      "error_name": "access_token_required",
      "description": "No access_token was passed"
    },
    {
      "error_id": 402,
      "error_name": "invalid_access_token",
      "description": "An access_token that is malformed, expired, or otherwise incorrect was passed"
    }
  ]
}
`
