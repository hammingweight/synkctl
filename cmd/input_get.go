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

// Reads the state of the input (e.g. solar panels) feeding the inverter
func readInputState(ctx context.Context) error {
	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	input, err := synkClient.Input(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInputState, err)
	}
	return displayObject(input.SynkObject)
}

// The input command allows an operator to get the state of the inputs feeding the inverter.
var inputGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets the state of the inverter's inputs",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return readInputState(cmd.Context())
	},
}

func init() {
	inputCmd.AddCommand(inputGetCmd)
	addKeysFlag(inputGetCmd)
}
