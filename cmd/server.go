package cmd

import (
	"github.com/jee-lee/budget-core/internal/category/repository"
	"github.com/jee-lee/budget-core/internal/category/server"
	"github.com/jee-lee/budget-core/internal/config"
	"github.com/jee-lee/budget-core/rpc/category"
	"github.com/spf13/cobra"
	"net/http"
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
	r := repository.NewRepository(*config.GetDB())
	server := server.NewServer(r)
	twirpHandler := category.NewCategoryServiceServer(server)
	err := http.ListenAndServe(":8080", twirpHandler)
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
}
