/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:     "input",
	Short:   "The inverter's input (e.g. solar panels, turbine)",
	Aliases: []string{"panels", "pv", "in", "inputs"},
}

func init() {
	rootCmd.AddCommand(inputCmd)
}
