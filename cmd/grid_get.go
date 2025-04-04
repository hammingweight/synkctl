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

// Reads the state of the grid (power being drawn from the grid, relay status, etc.)
func readGrid(ctx context.Context, sf pflag.Value) error {
	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	grid, err := synkClient.Grid(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadGridState, err)
	}
	if sf.String() == "true" {
		if len(keys) != 0 {
			return errors.New("cannot specify both \"--keys\" and \"--short\"")
		}
		keys = "acRealyStatus,etodayFrom,pac"
	}

	return displayObject(grid.SynkObject)
}

// The grid command allows an operator to get the grid's state
var gridGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the state of the grid connection",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return readGrid(cmd.Context(), cmd.Flags().Lookup("short").Value)
	},
}

func init() {
	gridCmd.AddCommand(gridGetCmd)
	addKeysFlag(gridGetCmd)
	gridGetCmd.Flags().BoolP("short", "s", false, "display short form of the grid state")
}
