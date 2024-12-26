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
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hammingweight/synkctl/configuration"
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
		Resp struct {
			Code int
			Body string
		}
		ExpectErr   bool
		AccessToken string
	}{{Name: "BadResponseCode",
		Resp: struct {
			Code int
			Body string
		}{404, ""},
		ExpectErr: true},
		{Name: "BadResponseBody",
			Resp: struct {
				Code int
				Body string
			}{200, "{this isn't json"},
			ExpectErr: true},
		{Name: "Unsuccessful",
			Resp: struct {
				Code int
				Body string
			}{200, `{"code":100, "msg":"this is a failure", "success":false, "data":{}}`},
			ExpectErr: true},
		{Name: "Successful",
			Resp: struct {
				Code int
				Body string
			}{200, `{"code":0, "msg":"", "success":true, "data":{"access_token":"12345"}}`},
			ExpectErr:   false,
			AccessToken: "12345"},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			url, cleanup := mockServer(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.Resp.Code)
				w.Write([]byte(tc.Resp.Body))
			})
			config := configuration.Configuration{Endpoint: url}
			client, err := Authenticate(context.Background(), &config)
			if tc.ExpectErr && err == nil {
				t.Error("Request should have failed")
			}
			if !tc.ExpectErr {
				if err != nil {
					t.Error("Request should have succeeded")
				}
				if tc.AccessToken != client.tokens.AccessToken {
					t.Errorf("Expected token %s, got %s\n", tc.AccessToken, client.tokens.AccessToken)
				}
			}
			defer cleanup()
		})
	}
}
