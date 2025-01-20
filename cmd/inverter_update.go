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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Updates the lower threshold battery capacity and/or the system work mode
func updateInverterSettings(ctx context.Context) error {
	essentialOnly := viper.GetString("essential-only")
	batteryCap := percentage(viper.GetString("battery-capacity"))
	gridCharge := viper.GetString("grid-charge")
	if essentialOnly == "" && batteryCap == "" && gridCharge == "" {
		return fmt.Errorf("%w: must supply \"essential-only\", \"battery-capacity\" or \"grid-charge\" flag",
			ErrCantUpdateInverterSettings)
	}

	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	inverterSettings, err := synkClient.Inverter(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInverterSettings, err)
	}

	if essentialOnly != "" {
		if err = inverterSettings.SetLimitedToLoad(essentialOnly == "on"); err != nil {
			return fmt.Errorf("%w: %w", ErrCantUpdateInverterSettings, err)
		}
	}

	if gridCharge != "" {
		inverterSettings.SetGridChargeOn(gridCharge == "on")
	}

	if batteryCap != "" {
		if err = inverterSettings.SetBatteryCapacity(batteryCap.Int()); err != nil {
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

	var batteryCap percentage
	updateCmd.Flags().VarP(&batteryCap, "battery-capacity", "b", "The minimum battery capacity")
	var essentialOnly onOff
	updateCmd.Flags().VarP(&essentialOnly, "essential-only", "e", "Power essential only (on) or all (off) circuits")
	var gridCharge onOff
	updateCmd.Flags().VarP(&gridCharge, "grid-charge", "g", "Enable (on) or disable (off) grid charging of the battery")

	viper.BindPFlag("battery-capacity", updateCmd.Flags().Lookup("battery-capacity"))
	viper.BindPFlag("essential-only", updateCmd.Flags().Lookup("essential-only"))
	viper.BindPFlag("grid-charge", updateCmd.Flags().Lookup("grid-charge"))
}
