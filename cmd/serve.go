package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zenzora/coveshare/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Serving on port 8080")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
