/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configurationCmd = &cobra.Command{
	Use:   "configuration",
	Short: "Configures credentials to access the SunSynk API",
	Long: `Commands to create a configuration file for accessing the SunSynk API and
to validate credentials.`,
}

func init() {
	rootCmd.AddCommand(configurationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
