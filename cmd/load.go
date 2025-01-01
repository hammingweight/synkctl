/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "The inverter's load statistics",
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
