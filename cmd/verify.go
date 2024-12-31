/*
Copyright © 2024 Carl Meijer
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/hammingweight/synkctl/synk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func verify(ctx context.Context, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("unexpected argument '%s'", args[0])
	}
	config := synk.Configuration{}
	err := config.ReadFromFile(viper.GetString("config"))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	_, err = synk.Authenticate(ctx, &config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	fmt.Println("OK.")
	return nil
}

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return verify(cmd.Context(), args)
	},
}

func init() {
	configurationCmd.AddCommand(verifyCmd)
}
