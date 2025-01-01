/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var batteryCmd = &cobra.Command{
	Use:   "battery",
	Short: "The inverter's battery state",
}

func init() {
	rootCmd.AddCommand(batteryCmd)
}
