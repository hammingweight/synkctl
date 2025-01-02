/*
Copyright © 2025 Carl Meijer
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var gridCmd = &cobra.Command{
	Use:   "grid",
	Short: "The state of the connection to the power grid",
}

func init() {
	rootCmd.AddCommand(gridCmd)
}
