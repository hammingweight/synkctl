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

// Lists all the inverter's that the user can view
func listInverters(ctx context.Context) error {
	synkClient, err := newClient(ctx, false)
	if err != nil {
		return nil
	}
	inverterSerialNumbers, err := synkClient.ListInverters(ctx)
	if err != nil {
		return err
	}
	marshalledBytes, err := json.MarshalIndent(inverterSerialNumbers, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(marshalledBytes))
	return nil
}

// The list command allows a user to get the serial number of inverters that
// they can manage
var listInverterCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all inverter serial numbers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listInverters(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(listInverterCmd)
}
