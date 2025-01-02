/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func readLoad(ctx context.Context) error {
	synkClient, err := getClient(ctx)
	if err != nil {
		return err
	}
	load, err := synkClient.ReadLoad(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadLoadStatistics, err)
	}
	return displayState(&load)
}

var getLoadCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets the inverter's current and cumulative load statistics",
	Aliases: []string{"read"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return readLoad(cmd.Context())
	},
}

func init() {
	loadCmd.AddCommand(getLoadCmd)
}
