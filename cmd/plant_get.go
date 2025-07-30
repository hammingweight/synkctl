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
	"errors"
	"fmt"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Reads the plant's status
func readPlant(ctx context.Context, sf pflag.Value, plantID int) error {
	if plantID == 0 {
		configFile := viper.GetString("config")
		config, err := configuration.ReadConfigurationFromFile(configFile)
		if err != nil {
			return err
		}
		plantID = config.DefaultPlantID
		if plantID == 0 {
			return fmt.Errorf("no plant ID specified")
		}
	}

	if sf.String() == "true" {
		if len(keys) != 0 {
			return errors.New("cannot specify both \"--keys\" and \"--short\"")
		}
		keys = "emonth,etoday,etotal,eyear,pac"
	}

	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	plant, err := synkClient.Plant(ctx, plantID)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadPlantStatus, err)
	}

	return displayObject(plant.SynkObject)
}

// The plant get command allows an operator to get the plant's status
var plantGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Reads the plant's status",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return readPlant(cmd.Context(), cmd.Flags().Lookup("short").Value, viper.GetInt("plant-id"))
	},
}

func init() {
	plantCmd.AddCommand(plantGetCmd)
	addKeysFlag(plantGetCmd)
	plantGetCmd.Flags().BoolP("short", "s", false, "get short output form")
	plantGetCmd.Flags().IntP("plant-id", "p", 0, "plant identifier")
	viper.BindPFlag("plant-id", plantGetCmd.Flag("plant-id"))
}
