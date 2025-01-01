/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "the inverter's load power state",
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
