/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func readInverterSettings(ctx context.Context) error {
	synkClient, err := getClient(ctx)
	if err != nil {
		return err
	}
	inverterSettings, err := synkClient.ReadInverterSettings(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInverterSettings, err)
	}
	return displayState(&inverterSettings)
}

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Reads the inverter settings",
	Aliases: []string{"read"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return readInverterSettings(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(getCmd)
}
