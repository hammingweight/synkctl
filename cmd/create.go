/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/hammingweight/synkctl/synk"
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
		user := viper.GetString("user")
		if user == "" {
			return errors.New("a user name (--user) must be supplied")
		}
		configFile := viper.GetString("config")
		password := viper.GetString("password")
		if password == "" {
			fmt.Fprintf(os.Stderr, "No password (--password) was supplied; you'll need to edit '%s' to add one.\n", configFile)
		}
		config := &synk.Configuration{
			Endpoint: viper.GetString("endpoint"),
			User:     user,
			Password: password,
		}
		err := config.WriteToFile(viper.GetString("config"))
		if err != nil {
			return err
		}
		fmt.Printf("Wrote configuration to '%s'.\n", configFile)
		return nil
	},
}

func init() {
	configurationCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("endpoint", "e", "https://api.sunsynk.net", "SunSynk API endpoint")
	createCmd.Flags().StringP("user", "u", "", "SunSynk user")
	createCmd.Flags().StringP("password", "p", "", "SunSynk user's password")

	// Viper bindings.
	viper.BindPFlag("endpoint", createCmd.Flags().Lookup("endpoint"))
	viper.BindPFlag("user", createCmd.Flags().Lookup("user"))
	viper.BindPFlag("password", createCmd.Flags().Lookup("password"))
}
