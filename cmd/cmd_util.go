package cmd

import (
	"fmt"
	"strings"

	"github.com/hammingweight/synkctl/rest"
	"github.com/spf13/viper"
)

func displayObject(o *rest.SynkObject) error {
	if viper.GetString("keys") != "" {
		var err error
		o, err = o.ExtractKeys(strings.Split(viper.GetString("keys"), ","))
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	fmt.Println(o)
	return nil
}
