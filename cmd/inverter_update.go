/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"context"
	"fmt"
	"strconv"

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

	if batteryCap != "" {
		_, err := strconv.Atoi(batteryCap)
		if err != nil {
			return fmt.Errorf("%w: battery-capacity must be an integer, not \"%s\"", ErrCantUpdateInverterSettings, batteryCap)
		}
		batteryCapUpper, ok := inverterSettings["batteryCap"]
		if !ok {
			return fmt.Errorf("%w: can't read upper limit for battery SOC", ErrCantUpdateInverterSettings)
		}
		batteryCapUpperInt, _ := strconv.Atoi(batteryCapUpper.(string))
		batteryCapInt, _ := strconv.Atoi(batteryCap)
		if batteryCapInt > batteryCapUpperInt {
			return fmt.Errorf("%w: \"battery-capacity\" cannot be greater than %d", ErrCantUpdateInverterSettings, batteryCapUpperInt)
		}
		batteryCapLower, ok := inverterSettings["batteryShutdownCap"]
		if !ok {
			return fmt.Errorf("%w: can't read lower limit for battery SOC", ErrCantUpdateInverterSettings)
		}
		batteryCapLowerInt, _ := strconv.Atoi(batteryCapLower.(string))
		if batteryCapInt <= batteryCapLowerInt {
			return fmt.Errorf("%w: \"battery-capacity\" must be greater than %d", ErrCantUpdateInverterSettings, batteryCapLowerInt)
		}
		_, ok = inverterSettings["cap7"]
		if ok {
			return fmt.Errorf("%w: more than six battery SOC settings", ErrCantUpdateInverterSettings)
		}
		for i := 1; i <= 6; i++ {
			key := fmt.Sprintf("cap%d", i)
			_, ok = inverterSettings[key]
			if !ok {
				return fmt.Errorf("%w: can't update setting \"%s\"", ErrCantUpdateInverterSettings, key)
			}
			inverterSettings[key] = batteryCap
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
