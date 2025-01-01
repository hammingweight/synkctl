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

func readInverterSettings(ctx context.Context) error {
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
	inverterSettings, err := synk.ReadInverterSettings(ctx, tokens, config.Endpoint, inverterSn)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInverterSettings, err)
	}
	settingsBytes, err := json.MarshalIndent(inverterSettings, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(settingsBytes))
	return nil
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Reads the inverter settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return readInverterSettings(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(getCmd)
}
