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

func readLoad(ctx context.Context) error {
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
	load, err := synk.ReadLoad(ctx, tokens, config.Endpoint, inverterSn)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantReadInverterSettings, err)
	}
	loadBytes, err := json.MarshalIndent(load, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(loadBytes))
	return nil
}

var getLoadCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the current and cumulative load",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return readLoad(cmd.Context())
	},
}

func init() {
	loadCmd.AddCommand(getLoadCmd)
}
