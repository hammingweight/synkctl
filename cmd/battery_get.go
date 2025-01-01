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

func readBattery(ctx context.Context) error {
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
	battery, err := synk.ReadBattery(ctx, tokens, config.Endpoint, inverterSn)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadBatteryState, err)
	}
	batteryBytes, err := json.MarshalIndent(battery, "", "    ")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadBatteryState, err)
	}
	fmt.Println(string(batteryBytes))
	return nil
}

var batteryGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Reads current and cumulative battery statistics",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return readBattery(cmd.Context())
	},
}

func init() {
	batteryCmd.AddCommand(batteryGetCmd)
}
