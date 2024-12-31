package synk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type SynkResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"msg"`
	Data    map[string]any `json:"data"`
	Success bool           `json:"success"`
}

func UmarshallResponseData(resp *http.Response, data any) error {
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
