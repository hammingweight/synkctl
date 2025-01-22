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

// Input is the SunSynk model of the inputs (e.g. solar panels) to an inverter.
// Attributes of the Grid include the number of inputs (e.g. two if there are
// two MPPTs) and the power being supplied by the input.
type Input struct{ *SynkObject }

// Input calls the SunSynk REST API to get the state of the input.
func (synkClient *SynkClient) Input(ctx context.Context) (*Input, error) {
	path := []string{"inverter", synkClient.SerialNumber, "realtime", "input"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readAPIV1(ctx, o, queryParams, path...)
	return &Input{o}, err
}

// Power returns the most recent reading of the power (in watts, W) being generated.
func (input *Input) Power() int {
	v, ok := input.Get("pac")
	if !ok {
		panic("cannot read the power being generated")
	}
	return int(v.(float64))
}

// PV returns the power, current and voltage from the n-th string where/
// the indexes start at 0.
func (input *Input) PV(n int) (map[string]any, bool) {
	v, ok := input.Get("pvIV")
	if !ok {
		panic("cannot read the pvIV attribute")
	}
	l, ok := v.([]any)
	if !ok || len(l) <= n {
		return nil, false
	}
	m := l[n]
	return m.(map[string]any), true
}
