package album

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearchAlbums(t *testing.T) {
	var params struct {
		Artist string `json:"artist_name"`
		Track  string `json:"track"`
	}
	t.Run("SUCCESS_FAIL", func(t *testing.T) {
		params.Artist = ""
		params.Track = ""
		data, _ := json.Marshal(&params)
		payload := strings.NewReader(string(data))
		req := httptest.NewRequest(http.MethodGet, "/", payload)
		rr := httptest.NewRecorder()
		SearchAlbums(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("wrong code: got %v want %v", status, http.StatusBadRequest)
		}
		if got, want := rr.Body.String(), `{"msg":"please input params."}`; got != want {
			t.Errorf("HelloWorld = %q, want %q", got, want)
		}
	})
	t.Run("SUCCESS_FULL", func(t *testing.T) {
		params.Artist = "Bodyslam"
		params.Track = "งมงาย"
		data, _ := json.Marshal(&params)
		payload := strings.NewReader(string(data))
		req := httptest.NewRequest(http.MethodGet, "/", payload)
		rr := httptest.NewRecorder()
		SearchAlbums(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
		if got, want := rr.Body.String(), "Bodyslam"; !strings.Contains(got, want) {
			t.Errorf("HelloWorld = %q, want %q", got, want)
		}

	})
}
