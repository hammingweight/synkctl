package synk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SynkResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	Data    map[string]any `json:"data"`
	Success bool           `json:"success"`
}

func umarshallResponseData(resp *http.Response, data any) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request returned status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	synkResponse := SynkResponse{}
	err = json.Unmarshal(body, &synkResponse)
	if err != nil {
		return err
	}
	if !synkResponse.Success {
		return errors.New(synkResponse.Message)
	}
	dataBytes, err := json.Marshal(synkResponse.Data)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(dataBytes, data)
	if err != nil {
		return nil
	}
	return nil
}

func readApiV1(ctx context.Context, tokens *Tokens, endpoint string, path ...string) (*http.Response, error) {
	fullPath := []string{"api", "v1"}
	fullPath = append(fullPath, path...)
	url, err := url.JoinPath(endpoint, fullPath...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokens.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}
