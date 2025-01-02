/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var configurationCmd = &cobra.Command{
	Use:     "configuration",
	Aliases: []string{"cfg", "config"},
	Short:   "Access configuration for the SunSynk API",
}

func init() {
	rootCmd.AddCommand(configurationCmd)
}
