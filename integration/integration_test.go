//go:build integration

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

package integration

import (
	"context"
	"os"
	"testing"
	"time"

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

	// Allow multiple attempts to authenticate, in case of flakiness somewhere.
	// Also, increase the delay between attempts.
	for i := 0; i < 12; i++ {
		time.Sleep(time.Duration(i) * time.Second * 5)
		client, err = rest.Authenticate(context.Background(), config)
		if err == nil {
			return
		}
	}
	panic(err)
}

func panicRecover(t *testing.T) {
	if v := recover(); v != nil {
		t.Fatal(v)
	}
}

func TestBattery(t *testing.T) {
	defer panicRecover(t)
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
	defer panicRecover(t)
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
	defer panicRecover(t)
	input, err := client.Input(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	// The next line will panic if we can't read the power.
	input.Power()
}

func TestInverter(t *testing.T) {
	defer panicRecover(t)
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
	defer panicRecover(t)
	load, err := client.Load(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	// The next line will panic if we can't read the power.
	load.Power()
}

func TestDetails(t *testing.T) {
	defer panicRecover(t)
	details, err := client.Details(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	power := details.RatedPower()
	if power < 1000 || power > 100000 {
		t.Errorf("rated power %d looks wrong", power)
	}
}
