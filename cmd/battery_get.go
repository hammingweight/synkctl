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
)

func readBattery(ctx context.Context) error {
	synkClient, err := getClient(ctx)
	if err != nil {
		return err
	}
	battery, err := synkClient.ReadBattery(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadBatteryState, err)
	}
	return displayState(&battery)
}

var batteryGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Reads battery statistics",
	Aliases: []string{"read"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return readBattery(cmd.Context())
	},
}

func init() {
	batteryCmd.AddCommand(batteryGetCmd)
}
