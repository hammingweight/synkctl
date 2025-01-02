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
	Aliases: []string{"ld"},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
