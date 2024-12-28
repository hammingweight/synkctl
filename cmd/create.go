/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("create called using config file %s\n", viper.GetString("config"))
		fmt.Println(viper.GetString("endpoint"))
	},
}

func init() {
	configurationCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("endpoint", "e", "https://api.sunsynk.net", "SunSynk API endpoint")
	createCmd.Flags().StringP("user", "u", "", "SunSynk username")
	createCmd.Flags().StringP("password", "p", "", "SunSynk user's password")
}
