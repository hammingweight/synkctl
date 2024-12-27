/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"os"

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

	// Flags that apply to all subcommands.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
