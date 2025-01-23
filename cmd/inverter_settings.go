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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func inverterSettings(ctx context.Context) error {
	synkClient, err := newClient(ctx, true)
	if err != nil {
		return err
	}
	settings, err := synkClient.Inverter(ctx)
	if err != nil {
		return err
	}
	is := settings.Settings()
	data, err := json.MarshalIndent(is, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

// settingsCmd represents the settings command
var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Summary of important inverter settings",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return inverterSettings(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(settingsCmd)
}
