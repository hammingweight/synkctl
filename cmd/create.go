/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("unexpected argument '%s'", args[0])
		}
		return createConfig(viper.GetString("endpoint"), viper.GetString("username"), viper.GetString("password"))
	},
}

func init() {
	configurationCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("endpoint", "e", "https://api.sunsynk.net", "SunSynk API endpoint")
	createCmd.Flags().StringP("username", "u", "", "SunSynk username")
	createCmd.Flags().StringP("password", "p", "", "SunSynk user's password")

	// Viper bindings.
	viper.BindPFlag("endpoint", createCmd.Flags().Lookup("endpoint"))
	viper.BindPFlag("username", createCmd.Flags().Lookup("username"))
	viper.BindPFlag("password", createCmd.Flags().Lookup("password"))
}

func createConfig(endpoint, username, password string) error {
	if username == "" {
		return errors.New("username must be specified")
	}

	return nil
}
