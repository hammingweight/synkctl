/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	// Flags that apply to all subcommands.
	// Get a default path to the synk config file.
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	configFile := filepath.Join(home, ".synk", "config")
	rootCmd.PersistentFlags().StringP("config", "c", configFile, "synkctl config file location")
	rootCmd.PersistentFlags().StringP("inverter", "i", "", "SunSynk inverter serial number")

	// Set up viper
	viper.SetEnvPrefix("SYNK")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("inverter", rootCmd.PersistentFlags().Lookup("inverter"))
}

// initConfig reads in ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
