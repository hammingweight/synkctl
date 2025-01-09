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

// Grid is the SunSynk model of the electricity grid connected to an inverter.
// Attributes of the Grid include the frequency (e.g. 50 Hz), voltage and,
// most importantly, whether the grid is up.
type Grid struct{ *SynkObject }

// Grid calls the SunSynk REST API to get the state of the power grid.
func (synkClient *SynkClient) Grid(ctx context.Context) (*Grid, error) {
	path := []string{"inverter", "grid", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber}
	o := &SynkObject{}
	err := synkClient.readAPIV1(ctx, o, queryParams, path...)
	return &Grid{o}, err
}

// IsUp returns true if the electricity grid can supply power. This is a convenience
// method that calls
//
//	grid.Get("acRealyStatus")
//
// to get the state of the grid (Note: "Realy" should probably be "Relay")
func (grid *Grid) IsUp() bool {
	v, ok := grid.Get("acRealyStatus")
	if !ok {
		panic("cannot determine whether the grid is up")
	}
	return int(v.(float64)) == 1
}
