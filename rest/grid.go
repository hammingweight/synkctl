package rest

import (
	"context"
	"errors"
)

// Grid is the SunSynk model of the electricity grid connected to an inverter.
// Attributes of the Grid include the frequency (e.g. 50 Hz), voltage and,
// most importantly, whether the grid is up.
type Grid struct{ *SynkObject }

// ReadBattery calls the SunSynk REST API to get the state of the power grid.
func (synkClient *SynkClient) ReadGrid(ctx context.Context) (*Grid, error) {
	path := []string{"inverter", "grid", synkClient.SerialNumber, "realtime"}
	queryParams := map[string]string{"sn": synkClient.SerialNumber}
	o := &SynkObject{}
	err := synkClient.readApiV1(ctx, o, queryParams, path...)
	return &Grid{o}, err
}

// IsUp returns true if the electricity grid can supply power. This is a convenience
// method that calls
//
//	grid.Get("acRealyStatus")
//
// to get the state of the grid (Note: "Realy" should probably be "Relay")
func (grid *Grid) IsUp() (bool, error) {
	v, ok := grid.Get("acRealyStatus")
	if !ok {
		return false, errors.New("cannot determine whether the grid is up")
	}
	return int(v.(float64)) == 1, nil
}
