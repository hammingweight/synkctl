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
	if battery.SOC() <= 0 || battery.SOC() > 100 {
		t.Errorf("battery SOC cannot be %d", battery.SOC())
	}
}
