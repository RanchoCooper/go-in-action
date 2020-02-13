package httpTest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("token", "a login token")
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestMockServer(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("start test mock server")

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("HTTP request failed")
	}

	defer resp.Body.Close()
	if resp.StatusCode != statusCode {
		fmt.Println(resp.Header.Get("token"))
		t.Fatalf("should receive \"%d\" status code. actual got: %v", statusCode, resp.StatusCode)
	}
	fmt.Println(resp.Header.Get("token"))
}
