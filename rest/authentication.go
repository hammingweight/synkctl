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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/hammingweight/synkctl/configuration"
)

type tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// SynkClient is a type that is needed for accessing the SunSynk API; it includes the API endpoint and OAuth
// tokens. It also includes the serial number for an inverter since most requests use the serial number to
// identify which object is being requested (e.g. the details of a battery connected to the inverter.)
type SynkClient struct {
	endpoint     string
	tokens       tokens
	SerialNumber string
}

type authenticationRequest struct {
	GrantType string `json:"grant_type"`
	User      string `json:"username"`
	Password  string `json:"password"`
}

func newAuthRequestBody(config *configuration.Configuration) (io.Reader, error) {
	authRequest := authenticationRequest{
		GrantType: "password",
		User:      config.User,
		Password:  config.Password,
	}
	r, err := json.Marshal(&authRequest)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(r), nil
}

// Authenticate uses a specify configuration to authenticate a user. If successful, a SynkClient is
// returned that can be used to make requests against the API.
func Authenticate(ctx context.Context, config *configuration.Configuration) (*SynkClient, error) {
	url, err := url.JoinPath(config.Endpoint, "oauth", "token")
	if err != nil {
		return nil, err
	}
	authRequest, err := newAuthRequestBody(config)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, authRequest)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	tokens := &tokens{}
	err = unmarshalResponseData(resp, tokens)
	if err != nil {
		return nil, err
	}
	return &SynkClient{endpoint: config.Endpoint, tokens: *tokens}, err
}
