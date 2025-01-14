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
	"fmt"
	"slices"
	"strconv"
)

// Inverter gets the settings of an inverter from the API.
func (synkClient *SynkClient) Inverter(ctx context.Context) (*Inverter, error) {
	path := []string{"common", "setting", synkClient.SerialNumber, "read"}
	inverter := &Inverter{}
	err := synkClient.readAPIV1(ctx, inverter, nil, path...)
	return inverter, err
}

// UpdateInverter issues a POST request to the SunSynk API to reconfigure the inverter. For
// example, the battery capacity can be changed or powering non-essential circuits can be disabled.
func (synkClient *SynkClient) UpdateInverter(ctx context.Context, settings *Inverter) error {
	path := []string{"common", "setting", synkClient.SerialNumber, "set"}
	postData, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	return updateAPIV1(ctx, synkClient, string(postData), path...)
}

func (synkClient *SynkClient) countInverters(ctx context.Context) (int, error) {
	path := []string{"inverters", "count"}
	resp := &map[string]any{}
	err := synkClient.readAPIV1(ctx, resp, nil, path...)
	return int((*resp)["total"].(float64)), err
}

const pageSize = 10

func (synkClient *SynkClient) inverterSerialNumbers(ctx context.Context, page int) ([]string, error) {
	path := []string{"inverters"}
	queryParams := map[string]string{}
	queryParams["page"] = strconv.Itoa(page)
	queryParams["limit"] = strconv.Itoa(pageSize)
	queryParams["type"] = "-2"
	queryParams["status"] = "-1"
	resp := &map[string]any{}
	err := synkClient.readAPIV1(ctx, resp, queryParams, path...)
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

// ListInverters returns the serial numbers of all inverters that the user can view.
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
		serialNumbers, err := synkClient.inverterSerialNumbers(ctx, i)
		if err != nil {
			return nil, err
		}
		inverterSerialNumbers = append(inverterSerialNumbers, serialNumbers...)
	}
	return inverterSerialNumbers, nil
}

// SetLimitedToLoad is a method that can disable (enable) power to non-essential loads.
// Passing the value 'true' will disable power from the inverter flowing to non-essential
// circuits (i.e. those powered by the CT coil).
func (settings *Inverter) SetLimitedToLoad(limitToLoad bool) error {
	sysWorkMode := settings.SysWorkMode
	if sysWorkMode != "1" && sysWorkMode != "2" {
		return fmt.Errorf("unexpected value for sysWorkMode setting: \"%s\"", sysWorkMode)
	}

	if limitToLoad {
		settings.SysWorkMode = "1"
	} else {
		settings.SysWorkMode = "2"
	}
	return nil
}

// LimitedToLoad returns true if the inverter powers only essential loads. If the inverter can
// power circuits connected to the CT, this method returns false.
func (settings *Inverter) LimitedToLoad() bool {
	if settings.SysWorkMode != "1" && settings.SysWorkMode != "2" {
		panic("unexpected value for sysWorkMode attribute: " + settings.SysWorkMode)
	}
	return settings.SysWorkMode == "1"
}

// SetBatteryCapacity sets the battery state of charge at which point the inverter will use the grid
// rather than batteries to power circuits. It is an error to set the capacity at a value higher than the
// maximum allowed SoC for the battery (typically, 100%) or below the shutdown capacity of the
// battery (e.g. 10%).
func (settings *Inverter) SetBatteryCapacity(batteryCap int) error {
	batteryCapUpperInt, _ := strconv.Atoi(settings.BatteryCap)
	if batteryCap > batteryCapUpperInt {
		return fmt.Errorf("\"battery-capacity\" cannot be greater than %d", batteryCapUpperInt)
	}

	batteryCapLowerInt, _ := strconv.Atoi(settings.BatteryShutdownCap)
	if batteryCap <= batteryCapLowerInt {
		return fmt.Errorf("\"battery-capacity\" must be greater than %d", batteryCapLowerInt)
	}

	batteryCapStr := fmt.Sprintf("%d", batteryCap)
	settings.Cap1 = batteryCapStr
	settings.Cap2 = batteryCapStr
	settings.Cap3 = batteryCapStr
	settings.Cap4 = batteryCapStr
	settings.Cap5 = batteryCapStr
	settings.Cap6 = batteryCapStr

	return nil
}

// BatteryCapacity gets the battery state of charge at which point the inverter will use the grid rather than batteries
// to power circuits.
func (settings *Inverter) BatteryCapacity() int {
	c := make([]int, 6)
	for i, s := range []string{settings.Cap1, settings.Cap2, settings.Cap3, settings.Cap4, settings.Cap5, settings.Cap6} {
		cc, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		c[i] = cc
	}
	return slices.Min(c)
}
