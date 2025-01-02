package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/hammingweight/synkctl/rest"
	"github.com/spf13/viper"
)

func getClient(ctx context.Context) (*rest.SynkClient, error) {
	configFile := viper.GetString("config")
	config, err := configuration.ReadConfigurationFromFile(configFile)
	if err != nil {
		return nil, err
	}
	inverterSn := viper.GetString("inverter")
	if inverterSn == "" {
		inverterSn = config.DefaultInverterSN
		if inverterSn == "" {
			return nil, ErrNoInverterSerialNumber
		}
	}
	synkClient, err := rest.Authenticate(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	synkClient.SerialNumber = inverterSn
	return synkClient, nil
}

func displayState(object *rest.SynkObject) error {
	objectBytes, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(objectBytes))
	return err
}
