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
	"errors"
	"strconv"
)

func (synkClient *SynkClient) countPlants(ctx context.Context, id int) (int, error) {
	path := []string{"user", strconv.Itoa(id), "plantCount"}
	resp := &map[string]any{}
	queryParams := map[string]string{"id": strconv.Itoa(id)}
	err := synkClient.readAPIV1(ctx, resp, queryParams, path...)
	if err != nil {
		return 0, err
	}
	return int((*resp)["total"].(float64)), nil
}

func (synkClient *SynkClient) plantIds(ctx context.Context, page int) ([]float64, error) {
	path := []string{"plants"}
	queryParams := map[string]string{}
	queryParams["page"] = strconv.Itoa(page)
	queryParams["limit"] = strconv.Itoa(pageSize)
	resp := &map[string]any{}
	if err := synkClient.readAPIV1(ctx, resp, queryParams, path...); err != nil {
		return nil, err
	}
	allPlants, ok := (*resp)["infos"]
	if !ok {
		return nil, errors.New("can't retrieve serial numbers from response")
	}
	plantList := allPlants.([]any)
	responseList := []float64{}
	for _, plant := range plantList {
		id, ok := plant.(map[string]any)["id"]
		if !ok {
			return nil, errors.New("can't retrieve plant numbers from response")
		}
		responseList = append(responseList, id.(float64))
	}
	return responseList, nil
}

// ListPlants returns the identifiers of all plants that the user can view.
func (synkClient *SynkClient) ListPlants(ctx context.Context, id int) ([]float64, error) {
	count, err := synkClient.countPlants(ctx, id)
	if err != nil {
		return nil, err
	}
	numPages := count / pageSize
	if count%pageSize != 0 {
		numPages++
	}
	plantIds := []float64{}
	for i := 1; i <= numPages; i++ {
		ids, err := synkClient.plantIds(ctx, i)
		if err != nil {
			return nil, err
		}
		plantIds = append(plantIds, ids...)
	}
	return plantIds, nil
}

// Plant is a model of an installation with an inverter, input and battery.
type Plant struct{ *SynkObject }

// Plant returns a plant object.
func (synkClient *SynkClient) Plant(ctx context.Context, plantID int) (*Plant, error) {
	path := []string{"plant", strconv.Itoa(plantID), "realtime"}
	queryParams := map[string]string{"plant": strconv.Itoa(plantID)}
	o := &SynkObject{}
	err := synkClient.readAPIV1(ctx, o, queryParams, path...)
	return &Plant{o}, err
}
