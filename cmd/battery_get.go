/*
Copyright © 2025 Carl Meijer
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
