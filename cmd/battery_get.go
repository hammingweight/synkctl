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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Reads the battery state (SoC, charging/discharging, etc)
func readBattery(ctx context.Context, sf pflag.Value) error {
	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	battery, err := synkClient.Battery(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadBatteryState, err)
	}
	if sf.String() == "true" {
		if len(keys) != 0 {
			return errors.New("cannot specify both \"--keys\" and \"--short\"")
		}
		keys = "bmsSoc,bmsVolt,emonthDischg,etodayDischg,etotalDischg,eyearDischg,power,temp,voltage"
	}

	return displayObject(battery.SynkObject)
}

// The battery command allows an operator to get the battery's state.
var batteryGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Reads battery statistics",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return readBattery(cmd.Context(), cmd.Flags().Lookup("short").Value)
	},
}

func init() {
	batteryCmd.AddCommand(batteryGetCmd)
	addKeysFlag(batteryGetCmd)
	batteryGetCmd.Flags().BoolP("short", "s", false, "display short form of the battery state")
}
