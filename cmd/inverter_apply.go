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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func confirmApply() bool {
	fmt.Println("You did not specify \"--force=true\". Do you really want to proceed?")
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
	var newSettings map[string]any
	err = json.Unmarshal(data, &newSettings)
	if err != nil {
		return err
	}

	settings, err := client.ReadInverterSettings(ctx)
	if err != nil {
		return err
	}
	for k, v := range newSettings {
		err = settings.Update(k, v)
		if err != nil {
			return err
		}
	}

	return client.UpdateInverterSettings(ctx, settings)
}

var inverterApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Applies the inverter settings from a file or stdin",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%v'", ErrUnexpectedArguments, args)
		}
		if !forceFlag {
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
}

var forceFlag bool

func init() {
	inverterCmd.AddCommand(inverterApplyCmd)
	inverterApplyCmd.Flags().BoolVar(&forceFlag, "force", false, "Acknowledge the risks")
	inverterApplyCmd.Flags().StringP("file", "f", "", "JSON file with inverter settings")
	viper.BindPFlag("file", inverterApplyCmd.Flags().Lookup("file"))
}
