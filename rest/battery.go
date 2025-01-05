package rest

import (
	"context"
	"errors"
)

// Battery is a model of a battery connected to an inverter.
type Battery struct{ *SynkObject }

// ReadBattery calls the SunSynk REST API to get the state of a battery connected
// to the inverter.
func (synkClient *SynkClient) ReadBattery(ctx context.Context) (*Battery, error) {
	path := []string{"inverter", "battery", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Battery{o}, err
}

// GetSOC returns the percentage state of charge of a battery. This is a convenience method that calls
//
//	battery.Get("bmsSoc")
//
// since "bmsSoc" is the attribute used by the SunSynk REST API.
func (battery *Battery) GetSOC() (int, error) {
	v, ok := battery.Get("bmsSoc")
	if ok {
		return int(v.(float64)), nil
	} else {
		return 0, errors.New("cannot retrieve the SOC")
	}
}
