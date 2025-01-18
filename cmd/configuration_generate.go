/*
Copyright 2025 Carl Meijer.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Writes the access credentials to a file
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
	if err := configuration.WriteConfigurationToFile(viper.GetString("config"), config); err != nil {
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

// The generate command creates a configuration file
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Creates a configuration file",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
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
