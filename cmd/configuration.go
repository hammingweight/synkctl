/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configurationCmd = &cobra.Command{
	Use:     "configuration",
	Aliases: []string{"cfg", "config"},
	Short:   "Access configuration for the SunSynk API",
	Long: `Commands to create a configuration file for accessing the SunSynk API and
to validate credentials.`,
}

func init() {
	rootCmd.AddCommand(configurationCmd)
}
