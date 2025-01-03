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

package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/hammingweight/synkctl/rest"
	"github.com/spf13/viper"
)

// Constructs a client to communicate with the SunSynk API
func newClient(ctx context.Context) (*rest.SynkClient, error) {
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

// Prints the state of a SunSynk object (battery, inverter, etc)
func displayState(object *rest.SynkObject) error {
	objectBytes, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		return err
	}
	_, err = fmt.Println(string(objectBytes))
	return err
}
