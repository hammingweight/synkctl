/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func readGrid(ctx context.Context) error {
	synkClient, err := getClient(ctx)
	if err != nil {
		return err
	}
	grid, err := synkClient.ReadGrid(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadGridState, err)
	}
	return displayState(&grid)
}

var gridGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets the state of the grid connection",
	Aliases: []string{"read"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return readGrid(cmd.Context())
	},
}

func init() {
	gridCmd.AddCommand(gridGetCmd)
}
