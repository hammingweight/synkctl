/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/hammingweight/synkctl/synk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func updateInverterSettings(ctx context.Context) error {
	essentialOnly := viper.GetString("essential-only")
	batteryCap := viper.GetString("battery-capacity")
	if essentialOnly == "" && batteryCap == "" {
		return fmt.Errorf("%w: must supply \"essential-only\" or \"battery-capacity\" flag", ErrCantUpdateInverterSettings)
	}
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
	if essentialOnly != "" {
		if inverterSettings["sysWorkMode"] != "1" && inverterSettings["sysWorkMode"] != "2" {
			return fmt.Errorf("%w: %s (%s)", ErrCantUpdateInverterSettings, "unexpected value for sysWorkMode setting: ", inverterSettings["sysWorkMode"])
		}
		switch essentialOnly {
		case "true":
			inverterSettings["sysWorkMode"] = "1"
		case "false":
			inverterSettings["sysWorkMode"] = "2"
		default:
			return fmt.Errorf("%w: essential-only must be \"true\" or \"false\", not \"%s\"", ErrCantUpdateInverterSettings, essentialOnly)
		}
	}
	return synk.UpdateInverterSettings(ctx, tokens, config.Endpoint, inverterSn, inverterSettings)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Basic options to update the inverter settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("%w '%s'", ErrUnexpectedArgument, args[0])
		}
		return updateInverterSettings(cmd.Context())
	},
}

func init() {
	inverterCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("battery-capacity", "b", "", "The minimum battery capacity")
	updateCmd.Flags().StringP("essential-only", "e", "", "Power essential only (true) or all (false) circuits")

	viper.BindPFlag("essential-only", updateCmd.Flags().Lookup("essential-only"))
	viper.BindPFlag("battery-capacity", updateCmd.Flags().Lookup("battery-capacity"))
}
