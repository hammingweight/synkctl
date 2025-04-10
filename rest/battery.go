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

// Battery is a model of a battery connected to an inverter.
type Battery struct{ *SynkObject }

// Battery calls the SunSynk REST API to get the state of a battery connected
// to the inverter.
func (synkClient *SynkClient) Battery(ctx context.Context) (*Battery, error) {
	path := []string{"inverter", "battery", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readAPIV1(ctx, o, queryParams, path...)
	return &Battery{o}, err
}

// SOC returns the percentage state of charge of a battery. This is a convenience method that calls
//
//	battery.Get("bmsSoc")
//
// since "bmsSoc" is the attribute used by the SunSynk REST API.
func (battery *Battery) SOC() (int, error) {
	v, ok := battery.Get("bmsSoc")
	if ok {
		switch v := v.(type) {
		case float64:
			return int(v), nil
		case int:
			return v, nil
		}
	}
	return 0, errors.New("cannot read battery SOC")
}

// Power returns the power being supplied from (positive value) or supplied to the battery. This is a convenience
// method that calls
//
//	battery.Get("power")
//
// since "power" is the attribute used by the SunSynk REST API.
func (battery *Battery) Power() (int, error) {
	v, ok := battery.Get("power")
	if ok {
		switch v := v.(type) {
		case float64:
			return int(v), nil
		}
	}
	return 0, errors.New("cannot read battery power")
}

// CapacityAh returns the capacity of the battery in ampere-hours.
func (battery *Battery) CapacityAh() (float64, error) {
	cap, ok := battery.SynkObject.Get("capacity")
	if ok {
		switch cap := cap.(type) {
		case string:
			v, err := strconv.ParseFloat(cap, 64)
			if err != nil {
				return 0.0, err
			}
			return v, nil
		case float64:
			return cap, nil
		}
	}
	return 0, errors.New("cannot read battery capacity")
}
