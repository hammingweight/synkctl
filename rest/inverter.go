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
	"encoding/json"
	"errors"
	"strconv"
)

func (synkClient *SynkClient) ReadInverterSettings(ctx context.Context) (*SynkObject, error) {
	path := []string{"common", "setting", synkClient.SerialNumber, "read"}
	return synkClient.readApiV1(ctx, nil, path...)
}

func (synkClient *SynkClient) UpdateInverterSettings(ctx context.Context, settings *SynkObject) error {
	path := []string{"common", "setting", synkClient.SerialNumber, "set"}
	postData, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	return updateApiV1(ctx, synkClient, string(postData), path...)
}

func (synkClient *SynkClient) countInverters(ctx context.Context) (int, error) {
	path := []string{"inverters", "count"}
	resp, err := synkClient.readApiV1(ctx, nil, path...)
	if err != nil {
		return 0, err
	}
	return int((*resp)["total"].(float64)), err
}

const pageSize = 10

func (synkClient *SynkClient) getInverterSerialNumbers(ctx context.Context, page int) ([]string, error) {
	path := []string{"inverters"}
	queryParams := map[string]string{}
	queryParams["page"] = strconv.Itoa(page)
	queryParams["limit"] = strconv.Itoa(pageSize)
	queryParams["type"] = "-2"
	queryParams["status"] = "-1"
	resp, err := synkClient.readApiV1(ctx, queryParams, path...)
	if err != nil {
		return nil, err
	}
	allInverters, ok := (*resp)["infos"]
	if !ok {
		return nil, errors.New("can't retrieve serial numbers from response")
	}
	inverterList := allInverters.([]any)
	responseList := []string{}
	for _, inv := range inverterList {
		sn, ok := inv.(map[string]any)["sn"]
		if !ok {
			return nil, errors.New("can't retrieve serial numbers from response")
		}
		responseList = append(responseList, sn.(string))
	}
	return responseList, nil
}

func (synkClient *SynkClient) ListInverters(ctx context.Context) ([]string, error) {
	count, err := synkClient.countInverters(ctx)
	if err != nil {
		return nil, err
	}
	numPages := count / pageSize
	if count%pageSize != 0 {
		numPages++
	}
	inverterSerialNumbers := []string{}
	for i := 1; i <= numPages; i++ {
		serialNumbers, err := synkClient.getInverterSerialNumbers(ctx, i)
		if err != nil {
			return nil, err
		}
		inverterSerialNumbers = append(inverterSerialNumbers, serialNumbers...)
	}
	return inverterSerialNumbers, nil
}
