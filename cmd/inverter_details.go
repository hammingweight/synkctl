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

// Reads the inverter's details.
func readDetails(ctx context.Context) error {
	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	details, err := synkClient.Details(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadDetails, err)
	}
	return displayObject(details.SynkObject)
}

// The details command returns the inverter's specification (like the rated power)
var detailsCmd = &cobra.Command{
	Use:     "details",
	Short:   "Returns the inverter's specification",
	Aliases: []string{"detail"},
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return readDetails(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(detailsCmd)
}
