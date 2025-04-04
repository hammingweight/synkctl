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
	"github.com/spf13/viper"
)

// Reads the inverter's settings (e.g. battery discharge threshold)
func readInverterSettings(ctx context.Context) error {
	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	inverterSettings, err := synkClient.Inverter(ctx)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInverterSettings, err)
	}
	if viper.GetBool("short") {
		if len(keys) != 0 {
			return errors.New("cannot specify both \"--keys\" and \"--short\"")
		}
		shortForm, err := inverterSettings.ToShortForm()
		if err != nil {
			return err
		}
		fmt.Println(shortForm)
		return nil
	}
	so, err := inverterSettings.ToSynkObject()
	if err != nil {
		return err
	}
	return displayObject(so)
}

// The inverter get command allows an operator to get the imverter's settings
var inverterGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "Reads the inverter settings",
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return readInverterSettings(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(inverterGetCmd)
	addKeysFlag(inverterGetCmd)

	inverterGetCmd.Flags().BoolP("short", "s", false, "Get short output (get only fields that can be updated)")
	viper.BindPFlag("short", inverterGetCmd.Flags().Lookup("short"))
}
