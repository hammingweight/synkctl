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
)

// Battery is a model of a battery connected to an inverter.
type Battery struct{ *SynkObject }

// Battery calls the SunSynk REST API to get the state of a battery connected
// to the inverter.
func (synkClient *SynkClient) Battery(ctx context.Context) (*Battery, error) {
	path := []string{"inverter", "battery", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Battery{o}, err
}

// SOC returns the percentage state of charge of a battery. This is a convenience method that calls
//
//	battery.Get("bmsSoc")
//
// since "bmsSoc" is the attribute used by the SunSynk REST API.
func (battery *Battery) SOC() int {
	v, ok := battery.Get("bmsSoc")
	if ok {
		return int(v.(float64))
	} else {
		panic("cannot retrieve the SOC")
	}
}
