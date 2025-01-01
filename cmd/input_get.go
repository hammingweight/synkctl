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

func readInputState(ctx context.Context) error {
	configFile := viper.GetString("config")
	config, err := synk.ReadConfigurationFromFile(configFile)
	if err != nil {
		return err
	}
	inverterSn := viper.GetString("inverter")
	if inverterSn == "" {
		inverterSn = config.DefaultInverterSN
		if inverterSn == "" {
			return ErrNoInverterSerialNumber
		}
	}
	tokens, err := synk.Authenticate(ctx, config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	input, err := synk.ReadInputState(ctx, tokens, config.Endpoint, inverterSn)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInputState, err)
	}
	inputBytes, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInputState, err)
	}
	fmt.Println(string(inputBytes))
	return nil
}

// getCmd represents the get command
var inputGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the state of the inverter's inputs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return readInputState(cmd.Context())
	},
}

func init() {
	inputCmd.AddCommand(inputGetCmd)
}
