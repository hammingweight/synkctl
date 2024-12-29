package synk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Tokens struct {
	Bearer  string
	Refresh string
}

type AuthenticationRequest struct {
	GrantType string `json:"grant_type"`
	User      string `json:"username"`
	Password  string `json:"password"`
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

func Authenticate(ctx context.Context, configFile string) (*Tokens, error) {
	config := &Configuration{}
	err := config.ReadFromFile(configFile)
	if err != nil {
		return nil, err
	}

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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication request returned status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return nil, err
}
