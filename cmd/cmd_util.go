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
	"fmt"
	"strings"

	"github.com/hammingweight/synkctl/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func displayObject(o *rest.SynkObject) error {
	if viper.GetString("keys") != "" {
		var err error
		o, err = o.ExtractKeys(strings.Split(viper.GetString("keys"), ","))
		if err != nil {
			return err
		}
	}
	fmt.Println(o)
	return nil
}

func addKeysFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("keys", "k", "", "Extract specific keys from response")
	viper.BindPFlag("keys", cmd.PersistentFlags().Lookup("keys"))
}
