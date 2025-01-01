/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Basic options to update the inverter settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("update called")
		return nil
	},
}

func init() {
	inverterCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("battery-capacity", "b", "", "The minimum battery capacity")
	updateCmd.Flags().StringP("essential-only", "e", "", "Power essential only (true) or all (false) circuits")
}
