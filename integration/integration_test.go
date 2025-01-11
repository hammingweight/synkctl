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
	// If we can't create a client, no tests will pass so
	// we just panic if something goes wrong.
	username := os.Getenv("TEST_USER")
	if username == "" {
		panic("the TEST_USER environment variable must be set")
	}
	password := os.Getenv("TEST_PASSWORD")
	if password == "" {
		panic("the TEST_PASSWORD environment variable must be set")
	}
	config, err := configuration.New(username, password)
	if err != nil {
		panic(err)
	}
	if os.Getenv("TEST_ENDPOINT") != "" {
		config.Endpoint = os.Getenv("TEST_ENDPOINT")
	}
	serialNumber := os.Getenv("TEST_INVERTER_SN")
	if serialNumber == "" {
		panic("the TEST_INVERTER_SN environment variable must be set")
	}
	config.DefaultInverterSN = serialNumber
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
	// The next line will panic if we can't read the power.
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
	// The next line will panic if we can't read the power.
	load.Power()
}
