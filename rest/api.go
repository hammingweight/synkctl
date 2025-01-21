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

// Package rest provides functions to query and update the SunSynk REST API.
package rest

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

// SynkResponse is the response returned by the SunSynk API to any HTTP GET. The "data" field encapsulates
// the object that was requested (e.g. OAuth tokens or details about an inverter, battery, etc.)
type SynkResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	Data    map[string]any `json:"data"`
	Success bool           `json:"success"`
}

func unmarshalResponseData(resp *http.Response, data any) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request returned status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	synkResponse := SynkResponse{}
	if err = json.Unmarshal(body, &synkResponse); err != nil {
		return err
	}
	if !synkResponse.Success {
		return errors.New(synkResponse.Message)
	}
	if data != nil {
		dataBytes, err := json.Marshal(synkResponse.Data)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(dataBytes, data); err != nil {
			return err
		}
	}
	return nil
}

func (synkClient *SynkClient) readAPIV1(ctx context.Context, synkObject any, queryParams map[string]string, path ...string) error {
	fullPath := []string{"api", "v1"}
	fullPath = append(fullPath, path...)
	url, err := url.JoinPath(synkClient.endpoint, fullPath...)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+synkClient.tokens.AccessToken)

	if queryParams != nil {
		q := req.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if err = unmarshalResponseData(resp, synkObject); err != nil {
		return err
	}
	return nil
}

func updateAPIV1(ctx context.Context, synkClient *SynkClient, contents string, path ...string) error {
	fullPath := []string{"api", "v1"}
	fullPath = append(fullPath, path...)
	url, err := url.JoinPath(synkClient.endpoint, fullPath...)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(contents))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+synkClient.tokens.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return unmarshalResponseData(resp, nil)
}
