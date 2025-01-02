/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// inverterCmd represents the inverter command
var inverterCmd = &cobra.Command{
	Use:     "inverter",
	Short:   "The inverter's settings",
	Aliases: []string{"inverters", "inv"},
}

func init() {
	rootCmd.AddCommand(inverterCmd)
}
