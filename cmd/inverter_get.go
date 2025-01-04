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

// Reads the inverter's settings (e.g. battery discharge threshold)
func readInverterSettings(ctx context.Context) error {
	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	inverterSettings, err := synkClient.ReadInverterSettings(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInverterSettings, err)
	}
	fmt.Println(inverterSettings)
	return nil
}

// The inverter command allows an operator to get/set the imverter's settings
var inverterGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Reads the inverter settings",
	Aliases: []string{"read"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%v'", ErrUnexpectedArguments, args)
		}
		return readInverterSettings(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(inverterGetCmd)
}
