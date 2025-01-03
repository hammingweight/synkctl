/*
Copyright 2025 Carl Meijer.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rest

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
