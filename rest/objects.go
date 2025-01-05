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
)

type SynkObject map[string]any

type Battery struct{ *SynkObject }

func (s SynkObject) String() string {
	m, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(m)
}

func (synkClient *SynkClient) ReadBattery(ctx context.Context) (*Battery, error) {
	path := []string{"inverter", "battery", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Battery{o}, err
}

type Grid struct{ *SynkObject }

func (synkClient *SynkClient) ReadGrid(ctx context.Context) (*Grid, error) {
	path := []string{"inverter", "grid", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Grid{o}, err
}

type InputState struct{ *SynkObject }

func (synkClient *SynkClient) ReadInputState(ctx context.Context) (*InputState, error) {
	path := []string{"inverter", synkClient.SerialNumber, "realtime", "input"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &InputState{o}, err
}

type Load struct{ *SynkObject }

func (synkClient *SynkClient) ReadLoad(ctx context.Context) (*Load, error) {
	path := []string{"inverter", "load", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Load{o}, err
}
