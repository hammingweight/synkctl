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
)

// Load is the SunSynk model of the load connected to an inverter.
// The most important attribute of the load is the power being consumed.
type Load struct{ *SynkObject }

// Load calls the SunSynk REST API to get the state of the load.
func (synkClient *SynkClient) Load(ctx context.Context) (*Load, error) {
	path := []string{"inverter", "load", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readAPIV1(ctx, o, queryParams, path...)
	return &Load{o}, err
}

// Power returns the current power (in watts, W) being consumed by the load.
// This is a convenience method that reads the totalPower attribute of a Load
// instance.
func (load *Load) Power() (int, error) {
	v, ok := load.Get("totalPower")
	// If the API is flaky, ok can be true but v is nil
	if ok {
		switch v := v.(type) {
		case float64:
			return int(v), nil
		}
	}
	return 0, errors.New("cannot determine the power being consumed")
}
