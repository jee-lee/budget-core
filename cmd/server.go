// See Makefile's `make server`
package cmd

import (
	"github.com/jee-lee/budget-core/internal/config"
	"github.com/jee-lee/budget-core/internal/repository"
	"github.com/jee-lee/budget-core/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		RunServer()
	},
}

func RunServer() {
	r := repository.NewRepository(config.GetDB())
	server.NewServer(r).Run()
}
