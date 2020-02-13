package handlers_test

import (
	"encoding/json"
	"github.com/RanchoCooper/go-in-action/chapter09/endpoint/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	handlers.Routes()
}

// TestSendJSON testing the sendjson internal endpoint
func TestSendJSON(t *testing.T) {

	req, err := http.NewRequest("GET", "/sendjson", nil)
	if err != nil {
		t.Fatal("\tfail to create a request.", ballotX, err)
	}
	t.Log("\tcreate a request", checkMark)

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	if rw.Code != http.StatusOK {
		t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
	}
	t.Log("\tShould receive \"200\"", checkMark)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		t.Fatal("\tShould decode the response.", ballotX)
	}
	t.Log("\tShould decode the response.", checkMark)

}
