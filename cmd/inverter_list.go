/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/hammingweight/synkctl/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func listInverters(ctx context.Context) error {
	configFile := viper.GetString("config")
	config, err := configuration.ReadConfigurationFromFile(configFile)
	if err != nil {
		return err
	}
	synkClient, err := rest.Authenticate(ctx, config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	inverterSerialNumbers, err := synkClient.ListInverters(ctx)
	if err != nil {
		return err
	}
	marshalledBytes, err := json.MarshalIndent(inverterSerialNumbers, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(marshalledBytes))
	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all inverter serial numbers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listInverters(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(listCmd)
}
