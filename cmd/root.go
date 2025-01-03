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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "synkctl",
	Short: "A CLI for SunSynk hybrid inverters",
	Long: `SynkCtl is a CLI for querying and updating SunSynk hybrid inverters and getting
the state of the battery, grid and input (e.g. solar panels) connected to the
inverter.`,
	Version: "0.2.0",
}

func Execute() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Get a default path to the synk config file.
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	configFile := filepath.Join(home, ".synk", "config")

	// Arguments that apply to all subcommands.
	rootCmd.PersistentFlags().StringP("config", "c", configFile, "synkctl config file location")
	rootCmd.PersistentFlags().StringP("inverter", "i", "", "SunSynk inverter serial number")

	// Set up viper
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("SYNK")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("inverter", rootCmd.PersistentFlags().Lookup("inverter"))
}

func initConfig() {
	// read in environment variables that match
	viper.AutomaticEnv()
}
