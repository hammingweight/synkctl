package synk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SynkResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	Data    map[string]any `json:"data"`
	Success bool           `json:"success"`
}

func unmarshallResponseData(resp *http.Response, data any) error {
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
	if data != nil {
		dataBytes, err := json.Marshal(synkResponse.Data)
		if err != nil {
			return nil
		}
		err = json.Unmarshal(dataBytes, data)
		if err != nil {
			return nil
		}
	}
	return nil
}

func readApiV1(ctx context.Context, tokens *Tokens, endpoint string, queryParams map[string]string, path ...string) (*http.Response, error) {
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

	if queryParams != nil {
		q := req.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

func updateApiV1(ctx context.Context, tokens *Tokens, endpoint string, contents string, path ...string) error {
	fullPath := []string{"api", "v1"}
	fullPath = append(fullPath, path...)
	url, err := url.JoinPath(endpoint, fullPath...)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(contents))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokens.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return unmarshallResponseData(resp, nil)
}
