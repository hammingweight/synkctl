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
	"strconv"
)

type Inverter struct{ *SynkObject }

func (synkClient *SynkClient) ReadInverterSettings(ctx context.Context) (*Inverter, error) {
	path := []string{"common", "setting", synkClient.SerialNumber, "read"}
	o, err := synkClient.readApiV1(ctx, nil, path...)
	if err != nil {
		return nil, err
	}
	return &Inverter{o}, nil
}

func (synkClient *SynkClient) UpdateInverterSettings(ctx context.Context, settings *Inverter) error {
	path := []string{"common", "setting", synkClient.SerialNumber, "set"}
	postData, err := json.Marshal(settings.SynkObject)
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

func (settings *Inverter) SetLimitToLoad(limitToLoad bool) error {
	sysWorkMode, ok := settings.Get("sysWorkMode")
	if !ok {
		return errors.New("could not get \"sysWorkMode\"")
	}
	if sysWorkMode != "1" && sysWorkMode != "2" {
		return fmt.Errorf("unexpected value for sysWorkMode setting: \"%s\"", sysWorkMode)
	}

	if limitToLoad {
		return settings.Update("sysWorkMode", "1")
	} else {
		return settings.Update("sysWorkMode", "2")
	}
}

func (settings *Inverter) GetLimitToLoad() (bool, error) {
	return false, nil
}

func (settings *Inverter) SetBatteryCapacity(batteryCap int) error {
	batteryCapUpper, ok := settings.Get("batteryCap")
	if !ok {
		return errors.New("can't read upper limit for battery SOC")
	}
	batteryCapUpperInt, _ := strconv.Atoi(batteryCapUpper.(string))
	if batteryCap > batteryCapUpperInt {
		return fmt.Errorf("\"battery-capacity\" cannot be greater than %d", batteryCapUpperInt)
	}
	batteryCapLower, ok := settings.Get("batteryShutdownCap")
	if !ok {
		return errors.New("can't read lower limit for battery SOC")
	}
	batteryCapLowerInt, _ := strconv.Atoi(batteryCapLower.(string))
	if batteryCap <= batteryCapLowerInt {
		return fmt.Errorf("\"battery-capacity\" must be greater than %d", batteryCapLowerInt)
	}
	for i := 1; i <= 6; i++ {
		key := fmt.Sprintf("cap%d", i)
		err := settings.Update(key, fmt.Sprintf("%d", batteryCap))
		if err != nil {
			return err
		}
	}
	_, ok = settings.Get("cap7")
	if ok {
		return errors.New("more than six battery SOC settings")
	}
	return nil
}

func (settings *Inverter) GetBatteryCapacity() (int, error) {
	cap1, ok := settings.Get("cap1")
	if !ok {
		return 0, errors.New("cannot get battery capacity \"cap1\"")
	}
	for i := 2; i <= 6; i++ {
		key := fmt.Sprintf("cap%d", i)
		if (*settings.SynkObject)[key] == cap1 {
			return 0, fmt.Errorf("battery capacity depends on the time of day")
		}
	}
	return strconv.Atoi(cap1.(string))
}
