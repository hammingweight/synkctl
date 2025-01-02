/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hammingweight/synkctl/synk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const pageSize = 10

func listInverters(ctx context.Context) error {
	configFile := viper.GetString("config")
	config, err := synk.ReadConfigurationFromFile(configFile)
	if err != nil {
		return err
	}
	tokens, err := synk.Authenticate(ctx, config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	count, err := synk.CountInverters(ctx, tokens, config.Endpoint)
	if err != nil {
		return err
	}
	numPages := count / pageSize
	if count%pageSize != 0 {
		numPages++
	}
	inverterSerialNumbers := []string{}
	for i := 1; i <= numPages; i++ {
		serialNumbers, err := synk.GetInverterSerialNumbers(ctx, tokens, config.Endpoint, i, pageSize)
		if err != nil {
			return err
		}
		inverterSerialNumbers = append(inverterSerialNumbers, serialNumbers...)
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
