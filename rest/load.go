package rest

import (
	"context"
	"errors"
)

// Load is the SunSynk model of the load connected to an inverter.
// The most important attribute of the load is the power being consumed.
type Load struct{ *SynkObject }

// ReadLoad calls the SunSynk REST API to get the state of the load.
func (synkClient *SynkClient) ReadLoad(ctx context.Context) (*Load, error) {
	path := []string{"inverter", "load", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Load{o}, err
}

func (load *Load) GetPower() (int, error) {
	v, ok := load.Get("totalPower")
	if !ok {
		return 0, errors.New("cannot determine the power being consumed")
	}
	return int(v.(float64)), nil
}
