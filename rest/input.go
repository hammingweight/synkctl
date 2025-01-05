package rest

import (
	"context"
	"errors"
)

// Input is the SunSynk model of the inputs (e.g. solar panels) to an inverter.
// Attributes of the Grid include the number of inputs (e.g. two if there are
// two MPPTs) and the power being supplied by the input.
type Input struct{ *SynkObject }

// ReadInput calls the SunSynk REST API to get the state of the input.
func (synkClient *SynkClient) ReadInputState(ctx context.Context) (*Input, error) {
	path := []string{"inverter", synkClient.SerialNumber, "realtime", "input"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber, "lan": "en"}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Input{o}, err
}

// GetPower returns the most recent reading of the power (in watts, W) being generated.
func (input *Input) GetPower() (int, error) {
	v, ok := input.Get("pac")
	if !ok {
		return 0, errors.New("cannot read the power being generated")
	}
	return int(v.(float64)), nil
}
