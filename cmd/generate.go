/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func generate() error {
	user := viper.GetString("user")
	if user == "" {
		return fmt.Errorf("%w: a user name (--user) must be supplied", ErrCantCreateConfigFile)
	}
	configFile := viper.GetString("config")
	password := viper.GetString("password")
	inverterSN := viper.GetString("inverter")
	config := &configuration.Configuration{
		Endpoint:          viper.GetString("endpoint"),
		User:              user,
		Password:          password,
		DefaultInverterSN: inverterSN,
	}
	err := configuration.WriteConfigurationToFile(viper.GetString("config"), config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantCreateConfigFile, err)
	}
	fmt.Printf("Wrote configuration to '%s'.\n", configFile)
	if password == "" {
		fmt.Fprintf(os.Stderr, "\nNo password (--password) was supplied; you'll need to edit \"%s\" to add one.\n", configFile)
	}
	if inverterSN == "" {
		fmt.Fprintf(os.Stderr, "\nNo inverter serial number (--inverter) was supplied, so no default serial number was written to the config file.\n")
	}
	return nil
}

// generateCmd represents the create command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Creates a configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return generate()
	},
}

func init() {
	configurationCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("endpoint", "e", configuration.DefaultEndpoint, "SunSynk API endpoint")
	generateCmd.Flags().StringP("user", "u", "", "SunSynk user")
	generateCmd.Flags().StringP("password", "p", "", "SunSynk user's password")

	// Viper bindings.
	viper.BindPFlag("endpoint", generateCmd.Flags().Lookup("endpoint"))
	viper.BindPFlag("user", generateCmd.Flags().Lookup("user"))
	viper.BindPFlag("password", generateCmd.Flags().Lookup("password"))
}
