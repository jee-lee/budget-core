package cmd

import (
	"github.com/jee-lee/budget-core/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		config.Logger.Fatal(err.Error())
		os.Exit(1)
	}
}
