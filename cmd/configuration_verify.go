/*
Copyright 2025 Carl Meijer.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"fmt"

	"github.com/hammingweight/synkctl/configuration"
	"github.com/hammingweight/synkctl/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Invoke the SunSynk OAUTH endpoint and check that the request succeeds
func verify(ctx context.Context, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("unexpected argument '%s'", args[0])
	}
	config, err := configuration.ReadConfigurationFromFile(viper.GetString("config"))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	_, err = rest.Authenticate(ctx, config)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCantAuthenticateUser, err)
	}
	fmt.Println("OK.")
	return nil
}

// The verify command checks that the access credentials work
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
