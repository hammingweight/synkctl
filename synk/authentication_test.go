package synk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer(h http.HandlerFunc) (string, func()) {
	ts := httptest.NewServer(h)

	return ts.URL, func() {
		ts.Close()
	}
}

func TestAuthenticate(t *testing.T) {
	tests := []struct {
		Name string
	}{{Name: "BadResponseCode"}}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			_, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {

			})
			defer cleanup()
		})
	}
}
