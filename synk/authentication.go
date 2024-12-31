package synk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

type AuthenticationRequest struct {
	GrantType string `json:"grant_type"`
	User      string `json:"username"`
	Password  string `json:"password"`
}

type AuthenticationResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    Tokens `json:"data"`
	Success bool   `json:"success"`
}

func getAuthRequestBody(config *Configuration) (io.Reader, error) {
	authRequest := AuthenticationRequest{
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

func Authenticate(ctx context.Context, config *Configuration) (*Tokens, error) {
	url, err := url.JoinPath(config.Endpoint, "oauth/token")
	if err != nil {
		return nil, err
	}
	authRequest, err := getAuthRequestBody(config)
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
	tokens := &Tokens{}
	err = UmarshallResponseData(resp, tokens)
	if err != nil {
		return nil, err
	}
	return tokens, err
}
