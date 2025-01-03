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

// Reads the current load from the inverter
func readLoad(ctx context.Context) error {
	synkClient, err := newClient(ctx)
	if err != nil {
		return err
	}
	load, err := synkClient.ReadLoad(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadLoadStatistics, err)
	}
	return displayState(&load)
}

// The load command displays load statistics
var getLoadCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets the inverter's current and cumulative load statistics",
	Aliases: []string{"read"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%v'", ErrUnexpectedArguments, args)
		}
		return readLoad(cmd.Context())
	},
}

func init() {
	loadCmd.AddCommand(getLoadCmd)
}
