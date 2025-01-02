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
	err = unmarshallResponseData(resp, tokens)
	if err != nil {
		return nil, err
	}
	return &SynkClient{endpoint: config.Endpoint, tokens: *tokens}, err
}
