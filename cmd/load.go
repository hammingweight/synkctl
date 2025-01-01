/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Reads the inverter's current and cumulative power load",
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
