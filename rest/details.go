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

// Details is a model of an inverter's specification.
type Details struct{ *SynkObject }

// Details calls the SunSynk REST API to get the inverter's specification.
func (synkClient *SynkClient) Details(ctx context.Context) (*Details, error) {
	path := []string{"inverter", synkClient.SerialNumber}
	o := &SynkObject{}
	err := synkClient.readAPIV1(ctx, o, nil, path...)
	return &Details{o}, err
}

// RatedPower returns the maximum power in watts that the inverter supports.
func (details *Details) RatedPower() (int, error) {
	power, ok := details.SynkObject.Get("ratePower")
	if ok {
		switch power := power.(type) {
		case float64:
			return int(power), nil
		case string:
			v, err := strconv.Atoi(power)
			if err != nil {
				return 0, err
			}
			return v, nil
		}
	}
	return 0, errors.New("cannot read rated power")
}
