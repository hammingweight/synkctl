/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:     "input",
	Short:   "The inverter's input (e.g. solar panels, turbine)",
	Aliases: []string{"panels", "pv"},
}

func init() {
	rootCmd.AddCommand(inputCmd)
}
