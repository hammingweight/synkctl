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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Updates the lower threshold battery capacity and/or the system work mode
func updateInverterSettings(ctx context.Context) error {
	essentialOnly := viper.GetString("essential-only")
	batteryCap := viper.GetString("battery-capacity")
	if essentialOnly == "" && batteryCap == "" {
		return fmt.Errorf("%w: must supply \"essential-only\" or \"battery-capacity\" flag", ErrCantUpdateInverterSettings)
	}

	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	inverterSettings, err := synkClient.Inverter(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInverterSettings, err)
	}

	// Check that we support only sysWorkModes "1" (power the essential loads only) and "2" (power all home circuits), i.e. we don't support
	// exporting to the grid
	if essentialOnly != "" {
		flag, err := strconv.ParseBool(essentialOnly)
		if err != nil {
			return fmt.Errorf("%w: essential-only must be \"true\" or \"false\", not \"%s\"", ErrCantUpdateInverterSettings, essentialOnly)
		}
		err = inverterSettings.SetLimitedToLoad(flag)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrCantUpdateInverterSettings, err)
		}
	}

	if batteryCap != "" {
		batteryCapInt, err := strconv.Atoi(batteryCap)
		if err != nil {
			return fmt.Errorf("%w: battery-capacity must be an integer, not \"%s\"", ErrCantUpdateInverterSettings, batteryCap)
		}
		err = inverterSettings.SetBatteryCapacity(batteryCapInt)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrCantUpdateInverterSettings, err)
		}
	}
	return synkClient.UpdateInverter(ctx, inverterSettings)
}

// Updates the configured lower discharge threshold for the bettery and whether to power all home circuits or
// only the essential circuits
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Basic options to update the inverter settings",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateInverterSettings(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("battery-capacity", "b", "", "The minimum battery capacity")
	updateCmd.Flags().StringP("essential-only", "e", "", "Power essential only (true) or all (false) circuits")

	viper.BindPFlag("essential-only", updateCmd.Flags().Lookup("essential-only"))
	viper.BindPFlag("battery-capacity", updateCmd.Flags().Lookup("battery-capacity"))
}
