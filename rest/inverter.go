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

	"github.com/hammingweight/synkctl/types"
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
	if err := synkClient.readAPIV1(ctx, resp, queryParams, path...); err != nil {
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

// SetEssentialOnly is an alias for the SetLimitedToLoad method.
func (settings *Inverter) SetEssentialOnly(essentialOnly bool) error {
	return settings.SetLimitedToLoad(essentialOnly)
}

// LimitedToLoad returns true if the inverter powers only essential loads. If the inverter can
// power circuits connected to the CT, this method returns false.
func (settings *Inverter) LimitedToLoad() bool {
	if settings.SysWorkMode != "1" && settings.SysWorkMode != "2" {
		panic("unexpected value for sysWorkMode attribute: " + settings.SysWorkMode)
	}
	return settings.SysWorkMode == "1"
}

// EssentialOnly is an alias for the LimitedToLoad method.
func (settings *Inverter) EssentialOnly() bool {
	return settings.LimitedToLoad()
}

// SetBatteryCapacity sets the battery state of charge at which point the inverter will use the grid
// rather than batteries to power circuits. It is an error to set the capacity at a value higher than the
// maximum allowed SoC for the battery (typically, 100%) or below the shutdown capacity of the
// battery (e.g. 10%).
func (settings *Inverter) SetBatteryCapacity(batteryCap int) error {
	if batteryCap > 100 {
		return fmt.Errorf("\"battery-capacity\" cannot be greater than 100")
	}

	batteryCapLow := settings.BatteryLowCapacity()
	if batteryCap <= batteryCapLow {
		return fmt.Errorf("\"battery-capacity\" must be greater than %d", batteryCapLow)
	}
	// The next check is for the pathological case where the shutdown capacity is greater than
	// the low SoC alarm setting.
	batteryCapShutdown := settings.BatteryShutdownCapacity()
	if batteryCap <= batteryCapShutdown {
		return fmt.Errorf("\"battery-capacity\" must be greater than %d", batteryCapShutdown)
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
func (settings *Inverter) BatteryCapacity() (int, error) {
	c := make([]int, 6)
	for i, s := range []string{settings.Cap1, settings.Cap2, settings.Cap3, settings.Cap4, settings.Cap5, settings.Cap6} {
		cc, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		c[i] = cc
	}
	return slices.Min(c), nil
}

// BatteryLowCapacity returns the battery state of charge that will generate an alarm.
func (settings *Inverter) BatteryLowCapacity() int {
	c, err := strconv.Atoi(settings.BatteryLowCap)
	if err != nil {
		panic(err)
	}
	return c
}

// BatteryShutdownCapacity returns the battery shutdown SoC.
func (settings *Inverter) BatteryShutdownCapacity() int {
	c, err := strconv.Atoi(settings.BatteryShutdownCap)
	if err != nil {
		panic(err)
	}
	return c
}

// GridChargeOn returns true if the grid is used at any time to charge the battery.
func (settings *Inverter) GridChargeOn() bool {
	gridCharge := []any{settings.Time1on, settings.Time2on, settings.Time3on, settings.Time4on, settings.Time5on, settings.Time6on}
	for _, c := range gridCharge {
		switch c := c.(type) {
		case string:
			b, err := strconv.ParseBool(c)
			if err != nil {
				panic(err)
			}
			if b {
				return b
			}
		case bool:
			if c {
				return c
			}
		default:
			panic(fmt.Sprintf("unexpected value for time-on setting: %v", c))
		}
	}
	return false
}

// SetGridChargeOn sets whether to enable (true) or disable (false) grid charging of the battery.
func (settings *Inverter) SetGridChargeOn(on bool) {
	settings.Time1on = on
	settings.Time2on = on
	settings.Time3on = on
	settings.Time4on = on
	settings.Time5on = on
	settings.Time6on = on
}

// Settings returns the settings for battery-capacity, essential-only and grid-charge.
func (settings *Inverter) Settings() (*InverterSettings, error) {
	is := &InverterSettings{}
	var err error
	is.BatteryCapacity, err = settings.BatteryCapacity()
	if err != nil {
		return nil, err
	}
	is.EssentialOnly = types.NewOnOff(settings.EssentialOnly())
	is.GridCharge = types.NewOnOff(settings.GridChargeOn())
	return is, nil
}

// SetSettings adjusts the battery-capacity, essential-only and grid-charge parameters
// of the inverter.
func (settings *Inverter) SetSettings(is *InverterSettings) error {
	settings.SetGridChargeOn(is.GridCharge.Bool())
	err := settings.SetBatteryCapacity(is.BatteryCapacity)
	if err != nil {
		return err
	}
	err = settings.SetEssentialOnly(is.EssentialOnly.Bool())
	return err
}
