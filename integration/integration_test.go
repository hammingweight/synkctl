//go:build integration

package integration

import (
	"context"
	"os"
	"testing"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/hammingweight/synkctl/rest"
)

var client *rest.SynkClient

func init() {
	username := os.Getenv("TEST_USER")
	password := os.Getenv("TEST_PASSWORD")
	config, err := configuration.New(username, password)
	if err != nil {
		panic(err)
	}
	config.DefaultInverterSN = os.Getenv("TEST_INVERTER_SN")
	client, err = rest.Authenticate(context.Background(), config)
	if err != nil {
		panic(err)
	}
}

func TestBattery(t *testing.T) {
	battery, err := client.Battery(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	soc := battery.SOC()
	if soc <= 0 || soc > 100 {
		t.Errorf("battery SOC cannot be %d", soc)
	}
}

func TestGrid(t *testing.T) {
	grid, err := client.Grid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	_, ok := grid.Get("fac")
	if !ok {
		t.Error("Cannot read the grid frequency")
	}
}

func TestInput(t *testing.T) {
	input, err := client.Input(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	input.Power()
}

func TestInverter(t *testing.T) {
	inverter, err := client.Inverter(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	batteryCap := inverter.BatteryCapacity()
	if batteryCap <= 0 || batteryCap > 100 {
		t.Errorf("battery capacity cannot be %d", batteryCap)
	}
}

func TestLoad(t *testing.T) {
	load, err := client.Load(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	load.Power()
}
