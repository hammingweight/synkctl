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
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/hammingweight/synkctl/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func confirmApply() bool {
	fmt.Println("You did not pass the argument \"force\". Do you really want to proceed?")
	fmt.Println("[yes/NO])")
	var resp string
	fmt.Scan(&resp)
	return resp == "yes"

}

func applyInverterSettings(ctx context.Context, in *os.File) error {
	client, err := newClient(ctx, true)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(in)
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	settings := &rest.Inverter{}
	if err = json.Unmarshal(data, settings); err != nil {
		return err
	}
	return client.UpdateInverter(ctx, settings)
}

// Updates the inverter settings from a file.
var inverterApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Applies the inverter settings from a file or stdin",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			if !confirmApply() {
				fmt.Println("aborting")
				os.Exit(1)
			}
		}
		filename := viper.GetString("file")
		in := os.Stdin
		if filename != "" {
			var err error
			in, err = os.Open(filename)
			if err != nil {
				return err
			}
		}
		return applyInverterSettings(cmd.Context(), in)
	},
	ValidArgs: []string{"force"},
}

func init() {
	inverterCmd.AddCommand(inverterApplyCmd)
	inverterApplyCmd.Flags().StringP("file", "f", "", "JSON file with inverter settings")
	viper.BindPFlag("file", inverterApplyCmd.Flags().Lookup("file"))
}
