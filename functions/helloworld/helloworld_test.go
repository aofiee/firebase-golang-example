package helloworld

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHTTP(t *testing.T) {
	type name struct {
		Name string
	}
	myname := name{
		Name: "Aofiee",
	}
	data, _ := json.Marshal(&myname)

	payload := strings.NewReader("")
	req := httptest.NewRequest(http.MethodGet, "/", payload)
	rr := httptest.NewRecorder()
	HelloHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}

	if got, want := rr.Body.String(), "Hello, World!"; got != want {
		t.Errorf("HelloWorld = %q, want %q", got, want)
	}
	payload = strings.NewReader(string(data))
	req = httptest.NewRequest(http.MethodGet, "/", payload)
	rr = httptest.NewRecorder()
	HelloHTTP(rr, req)
	if got, want := rr.Body.String(), "Hello, Aofiee!"; got != want {
		t.Errorf("HelloWorld = %q, want %q", got, want)
	}

	emptyName := name{
		Name: "",
	}
	dataEmpty, _ := json.Marshal(&emptyName)

	payload = strings.NewReader(string(dataEmpty))
	req = httptest.NewRequest(http.MethodGet, "/", payload)
	rr = httptest.NewRecorder()
	HelloHTTP(rr, req)
	if got, want := rr.Body.String(), "Hello, World!"; got != want {
		t.Errorf("HelloWorld = %q, want %q", got, want)
	}
}

func BenchmarkHelloHTTP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type name struct {
			Name string
		}
		emptyName := name{
			Name: "",
		}
		dataEmpty, _ := json.Marshal(&emptyName)
		payload := strings.NewReader(string(dataEmpty))
		req := httptest.NewRequest(http.MethodGet, "/", payload)
		rr := httptest.NewRecorder()
		HelloHTTP(rr, req)
	}
}
