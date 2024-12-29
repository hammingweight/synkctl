package synk

import (
	"context"
	"fmt"
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
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// TODO: Check status code
	fmt.Println(resp.StatusCode)
	return nil, err
}
